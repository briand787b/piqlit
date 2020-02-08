package controller

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/briand787b/piqlit/core/model"
	"github.com/briand787b/piqlit/core/obj"
	"github.com/briand787b/piqlit/core/perr"
	"github.com/briand787b/piqlit/core/plog"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// MediaController controls the flow of HTTP routes for Media resources
type MediaController struct {
	l  plog.Logger
	ms model.MediaTxCtlStore
	os obj.ObjectStore
}

// NewMediaController returns a new MediaController
func NewMediaController(l plog.Logger, ms model.MediaTxCtlStore, os obj.ObjectStore) *MediaController {
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
			render.Render(w, r, perr.NewInternalServerHTTPError(ctx, "no media_id in url params", c.l))
			return
		}

		mIDInt, err := strconv.Atoi(mID)
		if err != nil {
			render.Render(w, r, perr.NewValidationHTTPErrorFromError(ctx, err, "could not convert string mID to int", c.l))
			return
		}

		m, err := model.FindMediaByID(r.Context(), c.ms, mIDInt)
		if err != nil {
			render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not find Media by ID", c.l))
			return
		}

		ctx = context.WithValue(r.Context(), mediaCtxKey, m)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// HandleCreate handles Media creation
func (c *MediaController) HandleCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var data MediaRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, perr.NewValidationHTTPErrorFromError(ctx, err, "could not bind request body to Media", c.l))
		return
	}

	m := data.Media()
	if err := m.Persist(ctx, c.l, c.ms); err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not persist Media", c.l))
		return
	}

	// returned Media must be searched to get all created/updated children
	mr, err := model.FindMediaByID(ctx, c.ms, m.ID)
	if err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not find Media by ID", c.l))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewMediaResponse(mr))
}

// HandleDelete deletes the provided resource
func (c *MediaController) HandleDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	m, ok := ctx.Value(mediaCtxKey).(*model.Media)
	if !ok {
		render.Render(w, r, perr.NewNotFoundHTTPError(ctx, "no or invalid media value for mediaCtxKey", c.l))
		return
	}

	if err := m.Delete(ctx, c.l, c.ms); err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not delete resource", c.l))
		return
	}

	render.Status(r, http.StatusOK)
}

// HandleDownloadGZ x
func (c *MediaController) HandleDownloadGZ(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	m, ok := r.Context().Value(mediaCtxKey).(*model.Media)
	if !ok {
		render.Render(w, r, perr.NewNotFoundHTTPError(ctx, "no or invalid media value for mediaCtxKey", c.l))
		return
	}

	rc, err := m.DownloadGZ(ctx, c.l, c.os)
	if err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not get raw bytes", c.l))
		return
	}

	defer c.l.Close(ctx, rc)
	if _, err := io.Copy(w, rc); err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not send raw bytes", c.l))
		return
	}

	render.Status(r, http.StatusOK)
}

// HandleDownloadRaw x
func (c *MediaController) HandleDownloadRaw(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	m, ok := r.Context().Value(mediaCtxKey).(*model.Media)
	if !ok {
		render.Render(w, r, perr.NewNotFoundHTTPError(ctx, "no or invalid media value for mediaCtxKey", c.l))
		return
	}

	rc, err := m.DownloadRaw(ctx, c.l, c.os)
	if err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not get raw bytes", c.l))
		return
	}

	defer c.l.Close(ctx, rc)
	if _, err := io.Copy(w, rc); err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not send raw bytes", c.l))
		return
	}

	render.Status(r, http.StatusOK)
}

// HandleGetAllRoot writes a MediaResponse on the connection
func (c *MediaController) HandleGetAllRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ms, err := model.GetAllRootMedia(ctx, c.ms)
	if err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not get all root Media", c.l))
		return
	}

	render.Render(w, r, NewMediaResponseList(ms, 0, 0))
}

// HandleGetByID writes a MediaResponse on the connection
func (c *MediaController) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	m, ok := r.Context().Value(mediaCtxKey).(*model.Media)
	if !ok {
		render.Render(w, r, perr.NewNotFoundHTTPError(r.Context(), "no or invalid media value for mediaCtxKey", c.l))
		return
	}

	render.Render(w, r, NewMediaResponse(m))
}

// HandleUpdateShallow handles updates to the root Media specified in the request id
func (c *MediaController) HandleUpdateShallow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var data MediaUpdateRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not bind request body to Media", c.l))
		return
	}

	mBody := data.Media()

	if m, ok := ctx.Value(mediaCtxKey).(*model.Media); ok {
		mBody.ID = m.ID
	} else {
		render.Render(w, r, perr.NewNotFoundHTTPError(ctx, "no or invalid media value for mediaCtxKey", c.l))
		return
	}

	if err := mBody.UpdateShallow(ctx, c.l, c.ms); err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not update shallow", c.l))
		return
	}

	// returned Media must be searched to get all created/updated children
	mr, err := model.FindMediaByID(ctx, c.ms, mBody.ID)
	if err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not find Media by ID", c.l))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewMediaResponse(mr))
}

// HandleRawUpload handles the raw upload of a media file
func (c *MediaController) HandleRawUpload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	m, ok := ctx.Value(mediaCtxKey).(*model.Media)
	if !ok {
		render.Render(w, r, perr.NewNotFoundHTTPError(ctx, "media not found by id", c.l))
		return
	}

	defer c.l.Close(ctx, r.Body)
	m.Content = r.Body

	if err := m.UploadRaw(ctx, c.l, c.ms, c.os); err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not upload raw media", c.l))
		return
	}

	render.Status(r, http.StatusAccepted)
}

// HandleStreamRaw x
func (c *MediaController) HandleStreamRaw(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	m, ok := r.Context().Value(mediaCtxKey).(*model.Media)
	if !ok {
		render.Render(w, r, perr.NewNotFoundHTTPError(ctx, "no or invalid media value for mediaCtxKey", c.l))
		return
	}

	rc, err := m.DownloadRaw(ctx, c.l, c.os)
	if err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not get raw bytes", c.l))
		return
	}

	w.Header().Set("Content-Disposition", "inline")
	w.Header().Set("Content-Type", fmt.Sprintf(".%s", m.Encoding))

	defer c.l.Close(ctx, rc)
	if _, err := io.Copy(w, rc); err != nil {
		render.Render(w, r, perr.NewHTTPErrorFromError(ctx, err, "could not send raw bytes", c.l))
		return
	}
}
