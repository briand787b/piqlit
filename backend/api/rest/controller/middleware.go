package controller

import (
	"fmt"
	"net/http"

	"github.com/briand787b/piqlit/core/plog"
)

// Middleware acts as a bridge between the request and the controllers
type Middleware struct {
	l       plog.Logger
	uuidGen fmt.Stringer
}

// NewMiddleware returns a new Middleware
func NewMiddleware(l plog.Logger, uuidGen fmt.Stringer) *Middleware {
	return &Middleware{l: l, uuidGen: uuidGen}
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
