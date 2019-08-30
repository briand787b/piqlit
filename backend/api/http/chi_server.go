package phttp

import "github.com/go-chi/chi"

type ChiServer struct {
	s Server
	*chi.Mux
}

func NewChiServer(s Server) *ChiServer {
	cs := &ChiServer{
		s,
		chi.NewMux(),
	}

	cs.InitRouter()
	return cs
}

func (cs *ChiServer) InitRouter() {
}
