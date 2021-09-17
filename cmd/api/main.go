package main

import (
	"database/sql"
	"fmt"
	"log"
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

	e := routes.NewRouter(db)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
