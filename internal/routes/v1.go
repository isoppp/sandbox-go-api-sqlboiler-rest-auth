package routes

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func v1Routes(db *sql.DB) chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("v1 index"))
		if err != nil {
			log.Fatalln(err)
		}
	})
	return r
}
