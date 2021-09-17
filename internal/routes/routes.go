package routes

import (
	"database/sql"
	"time"

	"github.com/go-chi/render"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api", func(r chi.Router) {
		r.Mount("/", noVersioningRoutes(db))
		r.Mount("/v1", v1Routes(db))
	})
	return r
}
