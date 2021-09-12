package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/render"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Health struct {
	OK bool `json:"ok"`
}

func (h *Health) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api/", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			err := render.Render(w, r, &Health{OK: true})
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
		})
	})
	return r
}
