package controller

import (
	"fmt"
	"net/http"

	"github.com/briand787b/piqlit/core/obj"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

// Serve is a blocking function that serves HTTP
func Serve(port int, l *logrus.Logger, os obj.ObjectStore) {
	fc := FileController{l: l, os: os}

	// initialize router
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	// r.Use(mw.disableCORS)

	l.WithFields(logrus.Fields{
		"http_port": port,
	}).Info("http server has started")

	r.Get("/", fc.List)

	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
