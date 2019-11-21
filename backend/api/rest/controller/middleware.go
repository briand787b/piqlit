package controller

import (
	"fmt"
	"net/http"

	"github.com/briand787b/piqlit/core/plog"
)

const (
	// CorsEnvVarKey is the key used to get the CORS host
	CorsEnvVarKey = "PL_CORS_HOST"
)

// Middleware acts as a bridge between the request and the controllers
type Middleware struct {
	corsHost string
	l        plog.Logger
	uuidGen  fmt.Stringer
}

// NewMiddleware returns a new Middleware
func NewMiddleware(l plog.Logger, uuidGen fmt.Stringer, corsHost string) *Middleware {
	return &Middleware{l: l, uuidGen: uuidGen, corsHost: corsHost}
}

func (m *Middleware) spanAndTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := plog.StoreSpanIDTraceID(r.Context(), m.uuidGen.String(), m.uuidGen.String())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) logRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.l.Info(r.Context(), "started handling HTTP request", "uri", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) disableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", m.corsHost)
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Lang")

		next.ServeHTTP(w, r)
	})
}
