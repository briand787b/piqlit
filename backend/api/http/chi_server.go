package phttp

import (
	"github.com/briand787b/piqlit/api/http/controller"

	"github.com/go-chi/chi"
)

// ChiServer serves HTTP using the Chi router
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
	serverC := controller.NewNodeController(cs.ServerArgs.Logger, cs.ServerArgs.NodeStore)

	cs.Mux.Get("/", serverC.GetAllNodes)
}
