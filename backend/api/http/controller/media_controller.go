package controller

import (
	"net/http"

	plog "github.com/briand787b/piqlit/core/log"
	"github.com/briand787b/piqlit/core/model"
)

// MediaController controls the flow of HTTP routes for Media resources
type MediaController struct {
	l  plog.Logger
	ms model.MediaStore
}

// ListAllMedia lists all Media, bounded by the
func (mc *MediaController) ListAllMedia(w http.ResponseWriter, r *http.Request) {

}
