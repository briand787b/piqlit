package rest_test

import (
	"net/http"
	"testing"

	"github.com/briand787b/piqlit/api/rest"
)

func TestServerIsHTTPHandler(t *testing.T) {
	var _ http.Handler = rest.ChiServer{}
}
