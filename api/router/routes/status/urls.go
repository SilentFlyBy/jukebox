package status

import "github.com/go-chi/chi/v5"

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	addRoutes(r)

	return r
}

func addRoutes(r *chi.Mux) {
	r.Get("/health", GetHealthStatus)
}
