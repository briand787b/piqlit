package controller

import (
	"net/http"

	plog "github.com/briand787b/piqlit/core/log"
	"github.com/briand787b/piqlit/core/model"
)

// ServerController controls the flow of Server requests
type ServerController struct {
	l  plog.Logger
	ss model.ServerStore
}

// NewServerController returns a pointer to a new ServerController
func NewServerController(l plog.Logger, ss model.ServerStore) *ServerController {
	return &ServerController{l, ss}
}

// GetAllServers returns all the servers registered in the system
func (sc *ServerController) GetAllServers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, from ServerController"))
}
