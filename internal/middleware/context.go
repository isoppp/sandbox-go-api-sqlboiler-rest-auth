package middleware

import (
	"database/sql"
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"
	"sandbox-go-api-sqlboiler-rest-auth/models"

	"github.com/gorilla/securecookie"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type CustomContext struct {
	echo.Context
	Config       *config.Config
	ZapLogger    *zap.SugaredLogger
	DB           *sql.DB
	SecureCookie *securecookie.SecureCookie
	CurrentUser  *models.User
}

func NewCustomContext(c echo.Context, cfg *config.Config, l *zap.SugaredLogger, db *sql.DB, sc *securecookie.SecureCookie) *CustomContext {
	return &CustomContext{
		Context:      c,
		Config:       cfg,
		ZapLogger:    l,
		DB:           db,
		SecureCookie: sc,
		CurrentUser:  nil,
	}
}
