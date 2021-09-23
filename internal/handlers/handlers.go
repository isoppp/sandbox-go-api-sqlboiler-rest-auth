package handlers

import (
	"database/sql"
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"

	"github.com/gorilla/securecookie"

	"go.uber.org/zap"
)

type Handlers struct {
	cfg          *config.Config
	db           *sql.DB
	logger       *zap.SugaredLogger
	secureCookie *securecookie.SecureCookie
}

func NewHandler(cfg *config.Config, db *sql.DB, l *zap.Logger, sc *securecookie.SecureCookie) *Handlers {
	return &Handlers{
		cfg:          cfg,
		db:           db,
		logger:       l.Sugar(),
		secureCookie: sc,
	}
}
