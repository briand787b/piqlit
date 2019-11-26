package controller

import "net/http"

// VueController handles all vue-specific routes
type VueController struct{}

// NewVueController news up and returns a VueController
func NewVueController() *VueController {
	return &VueController{}
}

// HandleCorsPreflight handles preflight CORS checks for vue
func (c *VueController) HandleCorsPreflight(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
