package phttp_test

import (
	"net/http"
	"testing"

	phttp "github.com/briand787b/piqlit/api/http"
)

func TestServerIsHTTPHandler(t *testing.T) {
	var _ http.Handler = phttp.ChiServer{}
}
