package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"
	"sandbox-go-api-sqlboiler-rest-auth/internal/routes"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
		return
	}
	connStr := "host=localhost port=5433 dbname=sandbox user=postgres password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln("cannot open the database")
	}
	defer db.Close()

	r := routes.NewRouter(db)

	log.Printf("server is running on port %v", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r)
	if err != nil {
		log.Fatalln("cannot run server:", err)
		return
	}
}
