package http

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/golemcloud/golem-go/binding"
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
	var headerKeyValues []binding.WasiHttp0_2_0_TypesTuple2FieldKeyFieldValueT
	for key, values := range request.Header {
		for _, value := range values {
			headerKeyValues = append(headerKeyValues, binding.WasiHttp0_2_0_TypesTuple2FieldKeyFieldValueT{
				F0: key,
				F1: []byte(value),
			})
		}
	}
	headers := binding.StaticFieldsFromList(headerKeyValues).Unwrap()

	var method binding.WasiHttp0_2_0_TypesMethod
	switch strings.ToUpper(request.Method) {
	case "":
		method = binding.WasiHttp0_2_0_TypesMethodGet()
	case "GET":
		method = binding.WasiHttp0_2_0_TypesMethodGet()
	case "HEAD":
		method = binding.WasiHttp0_2_0_TypesMethodHead()
	case "POST":
		method = binding.WasiHttp0_2_0_TypesMethodPost()
	case "PUT":
		method = binding.WasiHttp0_2_0_TypesMethodPut()
	case "DELETE":
		method = binding.WasiHttp0_2_0_TypesMethodDelete()
	case "CONNECT":
		method = binding.WasiHttp0_2_0_TypesMethodConnect()
	case "OPTIONS":
		method = binding.WasiHttp0_2_0_TypesMethodOptions()
	case "TRACE":
		method = binding.WasiHttp0_2_0_TypesMethodTrace()
	case "PATCH":
		method = binding.WasiHttp0_2_0_TypesMethodPatch()
	default:
		method = binding.WasiHttp0_2_0_TypesMethodOther(request.Method)
	}

	path := request.URL.Path
	query := request.URL.RawQuery
	pathAndQuery := path
	if query != "" {
		pathAndQuery += "?" + query
	}

	var scheme binding.WasiHttp0_2_0_TypesScheme
	switch strings.ToLower(request.URL.Scheme) {
	case "http":
		scheme = binding.WasiHttp0_2_0_TypesSchemeHttp()
	case "https":
		scheme = binding.WasiHttp0_2_0_TypesSchemeHttps()
	default:
		scheme = binding.WasiHttp0_2_0_TypesSchemeOther(request.URL.Scheme)
	}

	userPassword := request.URL.User.String()
	var authority string
	if userPassword == "" {
		authority = request.URL.Host
	} else {
		authority = userPassword + "@" + request.URL.Host
	}

	requestHandle := binding.NewOutgoingRequest(headers)

	requestHandle.SetMethod(method)
	requestHandle.SetPathWithQuery(binding.Some(pathAndQuery))
	requestHandle.SetScheme(binding.Some(scheme))
	requestHandle.SetAuthority(binding.Some(authority))

	if request.Body != nil {
		reader := request.Body
		defer func() { _ = reader.Close() }()

		requestBodyResult := requestHandle.Body()
		if requestBodyResult.IsErr() {
			return nil, errors.New("failed to get request body")
		}
		requestBody := requestBodyResult.Unwrap()

		requestStreamResult := requestBody.Write()
		if requestStreamResult.IsErr() {
			return nil, errors.New("failed to start writing request body")
		}
		requestStream := requestStreamResult.Unwrap()

		buffer := make([]byte, 1024)
		for {
			n, err := reader.Read(buffer)

			result := requestStream.Write(buffer[:n])
			if result.IsErr() {
				requestStream.Drop()
				requestBody.Drop()
				return nil, errors.New("failed to write request body chunk")
			}

			if err == io.EOF {
				break
			}
		}

		requestStream.Drop()
		binding.StaticOutgoingBodyFinish(requestBody, binding.None[binding.WasiHttp0_2_0_TypesTrailers]())
		// requestBody.Drop() // TODO: this fails with "unknown handle index 0"
	}

	// TODO: timeouts
	connectTimeoutNanos := binding.None[uint64]()
	firstByteTimeoutNanos := binding.None[uint64]()
	betweenBytesTimeoutNanos := binding.None[uint64]()
	options := binding.NewRequestOptions()
	options.SetConnectTimeout(connectTimeoutNanos)
	options.SetFirstByteTimeout(firstByteTimeoutNanos)
	options.SetBetweenBytesTimeout(betweenBytesTimeoutNanos)

	futureResult := binding.WasiHttp0_2_0_OutgoingHandlerHandle(requestHandle, binding.Some(options))
	if futureResult.IsErr() {
		return nil, errors.New("failed to send request")
	}
	future := futureResult.Unwrap()

	incomingResponse, err := getIncomingResponse(future)
	if err != nil {
		return nil, err
	}

	status := incomingResponse.Status()
	responseHeaders := incomingResponse.Headers()
	defer responseHeaders.Drop()

	responseHeaderEntries := responseHeaders.Entries()
	header := http.Header{}

	for _, tuple := range responseHeaderEntries {
		ck := http.CanonicalHeaderKey(tuple.F0)
		header[ck] = append(header[ck], string(tuple.F1))
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

	responseBodyResult := incomingResponse.Consume()
	if responseBodyResult.IsErr() {
		return nil, errors.New("failed to consume response body")
	}
	responseBody := responseBodyResult.Unwrap()

	responseBodyStreamResult := responseBody.Stream()
	if responseBodyStreamResult.IsErr() {
		return nil, errors.New("failed to get response body stream")
	}
	responseBodyStream := responseBodyStreamResult.Unwrap()

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

func getIncomingResponse(future binding.WasiHttp0_2_0_OutgoingHandlerFutureIncomingResponse) (binding.WasiHttp0_2_0_TypesIncomingResponse, error) {
	result := future.Get()
	if result.IsSome() {
		result2 := result.Unwrap()
		if result2.IsErr() {
			return 0, errors.New("failed to send request")
		}
		result3 := result2.Unwrap()
		if result3.IsErr() {
			return 0, errors.New("failed to send request")
		}
		return result3.Unwrap(), nil
	} else {
		pollable := future.Subscribe()
		pollable.Block()
		return getIncomingResponse(future)
	}
}

type wasiStreamReader struct {
	Stream           binding.WasiHttp0_2_0_TypesInputStream
	Body             binding.WasiHttp0_2_0_TypesIncomingBody
	OutgoingRequest  binding.WasiHttp0_2_0_TypesOutgoingRequest
	IncomingResponse binding.WasiHttp0_2_0_TypesIncomingResponse
	Future           binding.WasiHttp0_2_0_TypesFutureIncomingResponse
}

func (reader *wasiStreamReader) Read(p []byte) (int, error) {
	c := cap(p)
	result := reader.Stream.BlockingRead(uint64(c))
	isEof := result.IsErr() && result.UnwrapErr() == binding.WasiIo0_2_0_StreamsStreamErrorClosed()
	if isEof {
		return 0, io.EOF
	} else if result.IsErr() {
		return 0, errors.New("failed to read response stream")
	} else {
		chunk := result.Unwrap()
		copy(p, chunk)
		return len(chunk), nil
	}
}

func (reader *wasiStreamReader) Close() error {
	reader.Stream.Drop()
	reader.Body.Drop()
	reader.IncomingResponse.Drop()
	reader.Future.Drop()
	reader.OutgoingRequest.Drop()
	return nil
}
