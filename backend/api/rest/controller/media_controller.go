package controller

import (
	"net/http"

	"github.com/briand787b/piqlit/core/model"
	"github.com/briand787b/piqlit/core/plog"
)

// MediaController controls the flow of HTTP routes for Media resources
type MediaController struct {
	l  plog.Logger
	ms model.MediaStore
}

// ListAllMedia lists all Media, bounded by the
func (mc *MediaController) ListAllMedia(w http.ResponseWriter, r *http.Request) {

}
