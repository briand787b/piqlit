package phttp

import "github.com/go-chi/chi"

type ChiServer struct {
	ServerAbstract
	*chi.Mux
}

func NewChiServer(s ServerAbstract) *ChiServer {
	cs := &ChiServer{
		s,
		chi.NewMux(),
	}

	cs.initRouter(s)
	return cs
}

func (cs *ChiServer) initRouter(s ServerAbstract) {
	cs.GetMediaByID(5)
}
