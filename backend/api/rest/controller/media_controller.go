package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/briand787b/piqlit/core/model"
	"github.com/briand787b/piqlit/core/obj"
	"github.com/briand787b/piqlit/core/plog"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

// MediaController controls the flow of HTTP routes for Media resources
type MediaController struct {
	l  plog.Logger
	ms model.MediaStore
	os obj.ObjectStore
}

// NewMediaController returns a new MediaController
func NewMediaController(l plog.Logger, ms model.MediaStore, os obj.ObjectStore) *MediaController {
	return &MediaController{
		l:  l,
		ms: ms,
		os: os,
	}
}

func (c *MediaController) mediaCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mID := chi.URLParam(r, "media_id")
		if mID == "" {
			render.Render(w, r, ErrInternalServer(tc.l, errors.New("could not get tag_id url param")))
			return
		}

		mIDInt, err := strconv.Atoi(mID)
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(tc.l, errors.Wrap(err, "could not convert string tag_id to int")))
			return
		}

		m, err := model.FindMediaByID(r.Context(), c.ms, mIDInt)
		if err != nil {
			if validation.IsValidationError(err) {
				render.Render(w, r, ErrNotFound(tc.l))
				return
			}

			render.Render(w, r, ErrInternalServer(tc.l, err))
			return
		}

		ctx := context.WithValue(r.Context(), mediaCtxKey, m)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetByID writes a MediaResponse on the connection
func (c *MediaController) GetByID(w http.ResponseWriter, r *http.Request) {

}

// // ListAllMedia lists all Media, bounded by the
// func (mc *MediaController) ListAllMedia(w http.ResponseWriter, r *http.Request) {

// }
