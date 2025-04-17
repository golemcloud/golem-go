package http

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	outgoinghandler "github.com/golemcloud/golem-go/binding/wasi/http/outgoing-handler"
	"github.com/golemcloud/golem-go/binding/wasi/http/types"
	"go.bytecodealliance.org/cm"
)

// WasiHttpTransport implements RoundTrip for the Golem WASI environment.
// It can be assigned to http.DefaultClient.Transport to globally set the default transport.
type WasiHttpTransport struct {
}

// InitStdDefaultClientTransport overrides the standard lib's default HTTP client's transport to the golem specific one
func InitStdDefaultClientTransport() {
	http.DefaultClient.Transport = &WasiHttpTransport{}
}

func (t *WasiHttpTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	var headerKeyValues []cm.Tuple[types.FieldKey, types.FieldValue]
	for key, values := range request.Header {
		for _, value := range values {
			headerKeyValues = append(headerKeyValues, cm.Tuple[types.FieldKey, types.FieldValue]{
				F0: types.FieldKey(key),
				F1: types.FieldValue(cm.ToList([]byte(value))),
			})
		}
	}
	headers, err, ok := types.FieldsFromList(cm.ToList(headerKeyValues)).Result()
	if !ok {
		return nil, errors.New(err.String())
	}

	var method types.Method
	switch strings.ToUpper(request.Method) {
	case "":
		method = types.MethodGet()
	case "GET":
		method = types.MethodGet()
	case "HEAD":
		method = types.MethodHead()
	case "POST":
		method = types.MethodPost()
	case "PUT":
		method = types.MethodPut()
	case "DELETE":
		method = types.MethodDelete()
	case "CONNECT":
		method = types.MethodConnect()
	case "OPTIONS":
		method = types.MethodOptions()
	case "TRACE":
		method = types.MethodTrace()
	case "PATCH":
		method = types.MethodPatch()
	default:
		method = types.MethodOther(request.Method)
	}

	path := request.URL.Path
	query := request.URL.RawQuery
	pathAndQuery := path
	if query != "" {
		pathAndQuery += "?" + query
	}

	var scheme types.Scheme
	switch strings.ToLower(request.URL.Scheme) {
	case "http":
		scheme = types.SchemeHTTP()
	case "https":
		scheme = types.SchemeHTTPS()
	default:
		scheme = types.SchemeOther(request.URL.Scheme)
	}

	userPassword := request.URL.User.String()
	var authority string
	if userPassword == "" {
		authority = request.URL.Host
	} else {
		authority = userPassword + "@" + request.URL.Host
	}

	requestHandle := types.NewOutgoingRequest(headers)

	requestHandle.SetMethod(method)
	requestHandle.SetPathWithQuery(cm.Some(pathAndQuery))
	requestHandle.SetScheme(cm.Some(scheme))
	requestHandle.SetAuthority(cm.Some(authority))

	if request.Body != nil {
		reader := request.Body
		defer func() { _ = reader.Close() }()

		requestBody, _, ok := requestHandle.Body().Result()
		if !ok {
			return nil, errors.New("failed to get request body")
		}

		requestStream, _, ok := requestBody.Write().Result()
		if !ok {
			return nil, errors.New("failed to start writing request body")
		}

		buffer := make([]byte, 1024)
		for {
			n, err := reader.Read(buffer)

			_, err2, ok := requestStream.Write(cm.ToList(buffer[:n])).Result()

			if !ok {
				requestStream.ResourceDrop()
				requestBody.ResourceDrop()
				return nil, errors.New(fmt.Sprintf("failed to write request body chunk: %s", err2.String()))
			}

			if err == io.EOF {
				break
			}
		}

		requestStream.ResourceDrop()

		types.OutgoingBodyFinish(requestBody, cm.None[types.Trailers]())
		// requestBody.Drop() // TODO: this fails with "unknown handle index 0"
	}

	// TODO: timeouts
	connectTimeoutNanos := cm.None[types.Duration]()
	firstByteTimeoutNanos := cm.None[types.Duration]()
	betweenBytesTimeoutNanos := cm.None[types.Duration]()
	options := types.NewRequestOptions()
	options.SetConnectTimeout(connectTimeoutNanos)
	options.SetFirstByteTimeout(firstByteTimeoutNanos)
	options.SetBetweenBytesTimeout(betweenBytesTimeoutNanos)

	future, err3, ok := outgoinghandler.Handle(requestHandle, cm.Some(options)).Result()
	if !ok {
		return nil, errors.New(fmt.Sprintf("failed to send request: %s", err3.String()))
	}

	incomingResponse, err4 := getIncomingResponse(future)
	if err4 != nil {
		return nil, err4
	}

	status := incomingResponse.Status()
	responseHeaders := incomingResponse.Headers()
	defer responseHeaders.ResourceDrop()

	responseHeaderEntries := responseHeaders.Entries()
	header := http.Header{}

	for _, tuple := range responseHeaderEntries.Slice() {
		ck := http.CanonicalHeaderKey(string(tuple.F0))
		header[ck] = append(header[ck], string([]byte(tuple.F1.Slice())))
	}

	var contentLength int64
	clHeader := header.Get("Content-Length")
	switch {
	case clHeader != "":
		cl, err := strconv.ParseInt(clHeader, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("net/http: ill-formed Content-Length header: %v", err)
		}
		if cl < 0 {
			// Content-Length values less than 0 are invalid.
			// See: https://datatracker.ietf.org/doc/html/rfc2616/#section-14.13
			return nil, fmt.Errorf("net/http: invalid Content-Length header: %q", clHeader)
		}
		contentLength = cl
	default:
		// If the response length is not declared, set it to -1.
		contentLength = -1
	}

	responseBody, _, ok := incomingResponse.Consume().Result()
	if !ok {
		return nil, errors.New("failed to consume response body")
	}

	responseBodyStream, _, ok := responseBody.Stream().Result()
	if !ok {
		return nil, errors.New("failed to get response body stream")
	}

	responseReader := wasiStreamReader{
		Stream:           responseBodyStream,
		Body:             responseBody,
		OutgoingRequest:  requestHandle,
		IncomingResponse: incomingResponse,
		Future:           future,
	}

	response := http.Response{
		Status:        fmt.Sprintf("%d %s", status, http.StatusText(int(status))),
		StatusCode:    int(status),
		Header:        header,
		ContentLength: contentLength,
		Body:          &responseReader,
		Request:       request,
	}

	return &response, nil
}

func getIncomingResponse(future types.FutureIncomingResponse) (types.IncomingResponse, error) {
	result := future.Get()
	result2 := result.Some()
	if result2 != nil {
		result3, _, ok := result2.Result()
		if !ok {
			return 0, errors.New("failed to send request")
		}
		result4, err, ok := result3.Result()
		if !ok {
			return 0, errors.New(fmt.Sprintf("failed to send request: %s", err.String()))
		}
		return result4, nil
	} else {
		pollable := future.Subscribe()
		pollable.Block()
		return getIncomingResponse(future)
	}
}

type wasiStreamReader struct {
	Stream           types.InputStream
	Body             types.IncomingBody
	OutgoingRequest  types.OutgoingRequest
	IncomingResponse types.IncomingResponse
	Future           types.FutureIncomingResponse
}

func (reader *wasiStreamReader) Read(p []byte) (int, error) {
	c := cap(p)
	chunk, err, ok := reader.Stream.BlockingRead(uint64(c)).Result()
	isEof := err.Closed()
	if isEof {
		return 0, io.EOF
	} else if !ok {
		return 0, errors.New(fmt.Sprintf("failed to read response stream: %s", err.String()))
	} else {
		copy(p, chunk.Slice())
		return len(chunk.Slice()), nil
	}
}

func (reader *wasiStreamReader) Close() error {
	reader.Stream.ResourceDrop()
	reader.Body.ResourceDrop()
	reader.IncomingResponse.ResourceDrop()
	reader.Future.ResourceDrop()
	reader.OutgoingRequest.ResourceDrop()
	return nil
}
