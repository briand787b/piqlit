package phttp

import (
	"github.com/briand787b/piqlit/api/http/controller"

	"github.com/go-chi/chi"
)

type ChiServer struct {
	*ServerArgs
	*chi.Mux
}

// NewChiServer instantiates a new ChiServer with initialized Chi router
func NewChiServer(isMaster bool, sa *ServerArgs) *ChiServer {
	cs := &ChiServer{
		sa,
		chi.NewMux(),
	}

	cs.initRouter(isMaster)
	return cs
}

func (cs *ChiServer) initRouter(isMaster bool) {
	serverC := controller.NewServerController(cs.ServerArgs.Logger, cs.ServerArgs.ServerStore)

	cs.Mux.Get("/", serverC.GetAllServers)
}
