package server

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"
	"sandbox-go-api-sqlboiler-rest-auth/internal/scookie"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewServer(cfg *config.Config, db *sql.DB, l *zap.Logger) *echo.Echo {
	sc := scookie.NewSecureCookie(cfg)
	e := echo.New()
	e.HideBanner = true
	e.Debug = cfg.IsDev

	bindGlobalMiddlewares(e, cfg, l, db, sc)
	bindRoutes(e, cfg, l, db, sc)

	if cfg.IsDev {
		exportRoutesJson(e)
	}

	return e
}

func exportRoutesJson(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile("routes.json", data, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
