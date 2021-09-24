package db

import (
	"database/sql"
	_ "embed"
	"log"
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"

	_ "github.com/lib/pq"
)

var (
	//go:embed seed.sql
	seedDoc string
)

func Seed(cfg *config.Config) error {
	// db
	db, err := sql.Open("postgres", cfg.GetDataSourceName())
	if err != nil {
		log.Fatalln("cannot open the database", cfg.GetDataSourceName())
	}
	defer func() {
		_ = db.Close()
	}()
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(seedDoc); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
