package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/briand787b/piqlit/core/model"
	"github.com/briand787b/piqlit/core/obj"
	"github.com/briand787b/piqlit/core/plog"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Serve is a blocking function that serves HTTP
func Serve(port int, l plog.Logger, ms model.MediaStore, os obj.ObjectStore) {
	mc := NewMediaController(l, ms, os)

	// initialize router
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/media", func(r chi.Router) {
		r.Post("/", mc.HandleCreate)

		r.Route("/{media_id}", func(r chi.Router) {
			r.With(mc.mediaCtx).Get("/", mc.HandleGetMediaByID)
		})
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}
