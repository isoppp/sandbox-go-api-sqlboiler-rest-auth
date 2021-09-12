package main

import (
	"fmt"
	"log"
	"net/http"
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"
	"sandbox-go-api-sqlboiler-rest-auth/internal/routes"
)

func main() {
	cfg, err := config.NewEnvConfig()
	if err != nil {
		log.Fatalln(err)
		return
	}

	r := routes.NewRouter()

	log.Printf("server is running on port %v", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r)
	if err != nil {
		log.Fatalln("cannot run server:", err)
		return
	}
}
