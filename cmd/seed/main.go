package main

import (
	"log"
	"sandbox-go-api-sqlboiler-rest-auth/db"
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"
)

func main() {
	cfg := config.NewConfig()
	err := db.Seed(cfg)
	if err != nil {
		log.Fatalln("cannot add seed data", err)
	}
	log.Println("seed data is added")
}
