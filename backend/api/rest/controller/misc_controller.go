package controller

import "net/http"

// MiscellaneousController handles miscellaneous routes that aren't specific to any other controller
type MiscellaneousController struct {
	defaultMethodNotAllowedHandler http.Handler
}

// NewMiscellaneousController returns a new MiscellaneousController
func NewMiscellaneousController(defaultMethodNotAllowedHandler http.Handler) *MiscellaneousController {
	return &MiscellaneousController{
		defaultMethodNotAllowedHandler: defaultMethodNotAllowedHandler,
	}
}

// HandleMethodNotAllowed overrides the default behavior of the router's MethodNotAllowed handler
func (c *MiscellaneousController) HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	c.defaultMethodNotAllowedHandler.ServeHTTP(w, r)
}
