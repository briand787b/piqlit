package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"

	"github.com/briand787b/piqlit/core/model"
	"github.com/briand787b/piqlit/core/obj"
	"github.com/briand787b/piqlit/core/plog"
	"github.com/google/uuid"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Serve is a blocking function that serves HTTP
func Serve(port int, l plog.Logger, ms model.MediaTxCtlStore, obs obj.ObjectStore) {
	vc := NewVueController()
	mc := NewMediaController(l, ms, obs)
	mw, err := NewMiddleware(l, uuid.New(), os.Getenv(CorsEnvVarKey))
	if err != nil {
		log.Fatalln(err)
	}

	// initialize router
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(mw.spanAndTrace)
	r.Use(mw.logRoute)
	r.Use(mw.disableCORS)

	r.Route("/media", func(r chi.Router) {
		r.Get("/", mc.HandleGetAllRoot)
		r.Post("/", mc.HandleCreate)
		r.Options("/", vc.HandleCorsPreflight)

		r.Route("/{media_id}", func(r chi.Router) {
			r.Options("/", vc.HandleCorsPreflight)
			r.With(mc.mediaCtx).Delete("/", mc.HandleDelete)
			r.With(mc.mediaCtx).Get("/", mc.HandleGetByID)
			r.With(mc.mediaCtx).Get("/download/gzip", mc.HandleDownloadGZ)
			r.With(mc.mediaCtx).Get("/download/raw", mc.HandleDownloadRaw)
			r.With(mc.mediaCtx).Put("/", mc.HandleUpdateShallow)
			r.With(mc.mediaCtx).Get("/stream/raw", mc.HandleStreamRaw)
			r.With(mc.mediaCtx).Put("/upload/raw", mc.HandleRawUpload)
		})
	})

	// // used by axios pre-flight checks in vue front end
	// r.Options("*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// }))

	ctx := plog.StoreSpanIDTraceID(context.Background(), "main", "main")
	walkFn := func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		funcName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
		l.Info(ctx, "API Route",
			"method", method,
			"route", route,
			"handler", funcName,
		)
		return nil
	}

	if err := chi.Walk(r, walkFn); err != nil {
		log.Fatalln("could not print API routes: ", err)
	}

	// TODO: delete me after merge
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}
