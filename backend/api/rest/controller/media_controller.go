package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/briand787b/piqlit/core/model"
	"github.com/briand787b/piqlit/core/obj"
	"github.com/briand787b/piqlit/core/perr"
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
		ctx := r.Context()
		mID := chi.URLParam(r, "media_id")
		if mID == "" {
			render.Render(w, r, newErrResponse(ctx, c.l, errors.New("no media_id in url params")))
			return
		}

		mIDInt, err := strconv.Atoi(mID)
		if err != nil {
			render.Render(w, r, newErrResponse(ctx, c.l, perr.NewErrInvalid("could not convert string tag_id to int")))
			return
		}

		m, err := model.FindMediaByID(r.Context(), c.ms, mIDInt)
		if err != nil {
			render.Render(w, r, newErrResponse(ctx, c.l, err))
			return
		}

		ctx = context.WithValue(r.Context(), mediaCtxKey, m)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// HandleCreate Handles Media creation
func (c *MediaController) HandleCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var data MediaRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, newErrResponse(ctx, c.l, perr.NewErrInvalid("could not bind request body to Media")))
		return
	}

	m := data.Media()
	if err := m.Persist(ctx, c.l, c.ms); err != nil {
		render.Render(w, r, newErrResponse(ctx, c.l, err))
		return
	}

	// returned Media must be searched to get all created/updated children
	mr, err := model.FindMediaByID(ctx, c.ms, m.ID)
	if err != nil {
		render.Render(w, r, newErrResponse(ctx, c.l, err))
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewMediaResponse(mr))
}

// HandleDelete deletes the provided resource
func (c *MediaController) HandleDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	m, ok := ctx.Value(mediaCtxKey).(*model.Media)
	if !ok {
		render.Render(w, r, newErrResponse(r.Context(), c.l, perr.NewErrNotFound(errors.New("no or invalid media value for mediaCtxKey"))))
		return
	}

	if err := m.Delete(ctx, c.ms); err != nil {
		render.Render(w, r, newErrResponse(ctx, c.l, errors.Wrap(err, "could not delete resource")))
	}

	render.Status(r, http.StatusOK)
}

// HandleGetByID writes a MediaResponse on the connection
func (c *MediaController) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	m, ok := r.Context().Value(mediaCtxKey).(*model.Media)
	if !ok {
		render.Render(w, r, newErrResponse(r.Context(), c.l, perr.NewErrNotFound(errors.New("no or invalid media value for mediaCtxKey"))))
		return
	}

	render.Render(w, r, NewMediaResponse(m))
}

// HandleUpdate handles updates to the root Media specified in the request id
func (c *MediaController) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	m, ok := ctx.Value(mediaCtxKey).(*model.Media)
	if !ok {
		render.Render(w, r, newErrResponse(r.Context(), c.l, perr.NewErrNotFound(errors.New("no or invalid media value for mediaCtxKey"))))
		return
	}

	var data MediaRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, newErrResponse(ctx, c.l, perr.NewErrInvalid("could not bind request body to Media")))
		return
	}

	mBody := data.Media()
	mBody.ID = m.ID
	if err := mBody.Persist(ctx, c.l, c.ms); err != nil {
		render.Render(w, r, newErrResponse(ctx, c.l, err))
		return
	}

	// returned Media must be searched to get all created/updated children
	mr, err := model.FindMediaByID(ctx, c.ms, m.ID)
	if err != nil {
		render.Render(w, r, newErrResponse(ctx, c.l, err))
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewMediaResponse(mr))
}
