package api

import (
	"net/http"

	"github.com/SilentFlyBy/jukebox/api/router/routes/cluster"
	"github.com/SilentFlyBy/jukebox/api/router/routes/status"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func createRouter() *chi.Mux {
	r := chi.NewRouter()
	addBasicMiddleware(r)
	addRoutes(r)

	return r
}

func addBasicMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
}

func addRoutes(r *chi.Mux) {
	r.Mount("/cluster", cluster.NewRouter())
	r.Mount("/status", status.NewRouter())
}

func StartApiServer() {
	r := createRouter()

	http.ListenAndServe(":3000", r)
}
