package api

import (
	"net/http"

	"github.com/SilentFlyBy/jukebox/base/id"
	"github.com/SilentFlyBy/jukebox/cluster"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type healthResponse struct {
	OK bool `json:"ok"`
}

func StartApiServer() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		health := healthResponse{OK: true}
		render.JSON(w, r, health)
	})

	r.Get("/cluster", func(w http.ResponseWriter, r *http.Request) {
		c := cluster.Cluster{
			ID:    id.NewID(),
			Label: "",
		}
		list := []cluster.Cluster{c}

		render.JSON(w, r, list)
	})

	http.ListenAndServe(":3000", r)
}
