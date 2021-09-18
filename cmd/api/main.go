package main

import (
	"database/sql"
	"fmt"
	"log"
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"
	"sandbox-go-api-sqlboiler-rest-auth/internal/routes"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

func main() {
	// config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
		return
	}

	// logger
	zapConfig := zap.NewDevelopmentConfig()
	encConfig := zap.NewDevelopmentEncoderConfig()
	encConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapConfig.EncoderConfig = encConfig
	logger, _ := zapConfig.Build()
	defer func() {
		_ = logger.Sync()
	}()

	// db
	connStr := "host=localhost port=5433 dbname=sandbox user=postgres password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln("cannot open the database")
	}
	defer func() {
		_ = db.Close()
	}()

	e := routes.NewRouter(db, logger)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
