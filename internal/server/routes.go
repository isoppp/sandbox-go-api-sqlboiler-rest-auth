package server

import (
	"database/sql"
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"
	"sandbox-go-api-sqlboiler-rest-auth/internal/handlers"

	"github.com/gorilla/securecookie"

	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
)

func bindRoutes(e *echo.Echo, cfg *config.Config, l *zap.Logger, db *sql.DB, sc *securecookie.SecureCookie) {
	// status
	e.GET("/api/status", handlers.GetStatus)

	// session
	e.POST("/api/v1/sessions", handlers.CreateSession)
	e.DELETE("/api/v1/sessions", handlers.DeleteSession)

	// users
	e.GET("/api/v1/me", handlers.Me)
	e.GET("/api/v1/users", handlers.GetUsers)
	e.POST("/api/v1/users", handlers.CreateUser)
	e.GET("/api/v1/users/:id", handlers.GetUser)
	e.PATCH("/api/v1/users/:id", handlers.PatchUser)
	e.DELETE("/api/v1/users/:id", handlers.DeleteUser)
}
