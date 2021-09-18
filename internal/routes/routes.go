package routes

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"sandbox-go-api-sqlboiler-rest-auth/internal/handlers"

	"go.uber.org/zap"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func NewRouter(db *sql.DB, logger *zap.Logger) *echo.Echo {
	e := echo.New()
	bindRouteMiddlewares(e, logger)
	// routes
	e.GET("/api/status", handlers.GetStatus)

	exportRoutesJson(e)
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

func bindRouteMiddlewares(e *echo.Echo, logger *zap.Logger) {
	// middlewares
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(ZapLogger(logger))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{}))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{}))

	// middlewares if production
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
	//	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//}))

	// middlewares if dev
	//e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	//	println(reqBody, resBody)
	//}))
}
