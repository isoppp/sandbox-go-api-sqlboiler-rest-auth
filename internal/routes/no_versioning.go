package routes

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func noVersioningRoutes(db *sql.DB) chi.Router {
	r := chi.NewRouter()
	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("working"))
		if err != nil {
			log.Fatalln(err)
		}
	})
	return r
}
