package server

import (
	"database/sql"
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"
	"sandbox-go-api-sqlboiler-rest-auth/internal/middleware"

	"github.com/gorilla/securecookie"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func bindGlobalMiddlewares(e *echo.Echo, cfg *config.Config, logger *zap.Logger, db *sql.DB, sc *securecookie.SecureCookie) {
	sl := logger.Sugar()
	// custom context
	e.Use(middleware.SetCustomContext(cfg, sl, db, sc))

	// middlewares
	e.Pre(echomiddleware.RemoveTrailingSlash())
	e.Use(middleware.RequestZapLogger(logger))
	e.Use(echomiddleware.RateLimiter(echomiddleware.NewRateLimiterMemoryStore(20)))
	e.Use(echomiddleware.SecureWithConfig(echomiddleware.SecureConfig{}))
	e.Use(echomiddleware.TimeoutWithConfig(echomiddleware.TimeoutConfig{}))
	e.Use(middleware.SessionRestorer(db, sl, sc))

	// middlewares if production
	if !cfg.IsDev {
		e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
			AllowOrigins: []string{"https://labstack.com", "https://labstack.net"}, // TODO set real url
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
	}

	// middlewares if dev
	if cfg.IsDev {
		e.Use(echomiddleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
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
