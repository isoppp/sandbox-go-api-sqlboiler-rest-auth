package handlers

import (
	"database/sql"

	"go.uber.org/zap"
)

type Handlers struct {
	db      *sql.DB
	logger  *zap.Logger
	slogger *zap.SugaredLogger
}

func NewHandler(db *sql.DB, l *zap.Logger) *Handlers {
	var handler Handlers
	handler.db = db
	handler.logger = l
	handler.slogger = l.Sugar()
	return &handler
}
