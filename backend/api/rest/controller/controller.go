package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/briand787b/piqlit/core/model"
	"github.com/briand787b/piqlit/core/obj"
	"github.com/briand787b/piqlit/core/plog"
	"github.com/google/uuid"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Serve is a blocking function that serves HTTP
func Serve(port int, l plog.Logger, ms model.MediaTxCtlStore, os obj.ObjectStore) {
	mc := NewMediaController(l, ms, os)
	mw := NewMiddleware(l, uuid.New())

	// initialize router
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(mw.spanAndTrace)

	r.Route("/media", func(r chi.Router) {
		r.Post("/", mc.HandleCreate)

		r.Route("/{media_id}", func(r chi.Router) {
			r.With(mc.mediaCtx).Delete("/", mc.HandleDelete)
			r.With(mc.mediaCtx).Get("/", mc.HandleGetByID)
			r.With(mc.mediaCtx).Get("/download/gzip", mc.HandleDownloadGZ)
			r.With(mc.mediaCtx).Get("/download/raw", mc.HandleDownloadRaw)
			r.With(mc.mediaCtx).Put("/", mc.HandleUpdateShallow)
			r.With(mc.mediaCtx).Get("/stream/raw", mc.HandleStreamRaw)
			r.With(mc.mediaCtx).Put("/upload/raw", mc.HandleRawUpload)
		})
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}
