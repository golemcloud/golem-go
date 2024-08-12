package http_test

import (
	"net/http"
	"testing"

	golemhttp "github.com/golemcloud/golem-go/net/http"
)

func TestRoundTripCompilesAndCanBeUsedAsDefaultClientTransport(*testing.T) {
	http.DefaultClient.Transport = &golemhttp.WasiHttpTransport{}
}
