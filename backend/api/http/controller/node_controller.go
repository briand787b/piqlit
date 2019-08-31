package controller

import (
	"net/http"

	plog "github.com/briand787b/piqlit/core/log"
	"github.com/briand787b/piqlit/core/model"
)

// NodeController controls the flow of Node requests
type NodeController struct {
	l  plog.Logger
	ns model.NodeStore
}

// NewNodeController returns an initialized NodeController
func NewNodeController(l plog.Logger, ns model.NodeStore) *NodeController {
	return &NodeController{
		l:  l,
		ns: ns,
	}
}

// GetAllNodes returns all the nodes registered in the system
func (sc *NodeController) GetAllNodes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, from NodeController"))
}
