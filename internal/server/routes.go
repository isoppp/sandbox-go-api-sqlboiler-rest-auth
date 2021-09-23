package server

import (
	"database/sql"
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"
	"sandbox-go-api-sqlboiler-rest-auth/internal/handlers"

	"github.com/gorilla/securecookie"

	"go.uber.org/zap"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func bindRoutes(e *echo.Echo, cfg *config.Config, l *zap.Logger, db *sql.DB, sc *securecookie.SecureCookie) {
	h := handlers.NewHandler(cfg, db, l, sc)
	bindGlobalMiddlewares(e, cfg, l, db, sc)

	// status
	e.GET("/api/status", h.GetStatus)

	// session
	e.POST("/api/v1/sessions", h.CreateSession)
	e.DELETE("/api/v1/sessions", h.DeleteSession)

	// users
	e.GET("/api/v1/users", h.GetUsers)
	e.POST("/api/v1/users", h.CreateUser)
	e.GET("/api/v1/users/:id", h.GetUser)
	e.PATCH("/api/v1/users/:id", h.PatchUser)
	e.DELETE("/api/v1/users/:id", h.DeleteUser)
}

func bindGlobalMiddlewares(e *echo.Echo, cfg *config.Config, logger *zap.Logger, db *sql.DB, sc *securecookie.SecureCookie) {
	sl := logger.Sugar()
	// middlewares
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(ZapLogger(logger))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{}))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{}))
	e.Use(SessionRestorer(db, sl, sc))

	// middlewares if production
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
	//	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//}))

	// middlewares if dev
	if cfg.IsDev {
		e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			if len(reqBody) == 0 {
				sl.Debug("request body: ", "None")
			} else {
				sl.Debug("request body: ", string(reqBody))
			}

			if len(resBody) == 0 {
				sl.Debug("response body: ", "No Content")
			} else {
				sl.Debug("response body: ", string(resBody))
			}
		}))
	}
}
