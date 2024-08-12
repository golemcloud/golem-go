package http

import (
	"net/http"
	"testing"
)

func TestRoundTripCompilesAndCanBeUsedAsDefaultClientTransport(*testing.T) {
	http.DefaultClient.Transport = &WasiHttpTransport{}
}
