package phttp

import (
	"net/http"

	plog "github.com/briand787b/piqlit/core/log"
	"github.com/briand787b/piqlit/core/model"
)

// Server x
type Server interface {
	plog.Logger
	model.MediaStore
	model.ServerStore
	http.Handler
}

// func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	s.Mux.ServeHTTP(w, r)
// }
