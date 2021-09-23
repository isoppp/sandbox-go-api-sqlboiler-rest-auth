package main

import (
	"database/sql"
	"fmt"
	"log"
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"
	"sandbox-go-api-sqlboiler-rest-auth/internal/logger"
	"sandbox-go-api-sqlboiler-rest-auth/internal/server"

	"github.com/volatiletech/sqlboiler/v4/boil"

	_ "github.com/lib/pq"
)

func main() {
	// config
	cfg := config.NewConfig()

	// db
	db, err := sql.Open("postgres", cfg.GetDataSourceName())
	if err != nil {
		log.Fatalln("cannot open the database", cfg.GetDataSourceName())
	}
	defer func() {
		_ = db.Close()
	}()

	// enable debug mode
	boil.DebugMode = cfg.IsDev

	// logger
	zl, err := logger.NewLogger(cfg)
	if err != nil {
		log.Fatalln("cannot create logger")
	}
	defer func() {
		_ = zl.Sync()
	}()

	// router
	e := server.NewServer(cfg, db, zl)

	err = e.Start(fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		zl.Sugar().Fatal(err)
	}
}
