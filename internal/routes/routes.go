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

func NewRouter(db *sql.DB, l *zap.Logger) *echo.Echo {
	h := handlers.NewHandler(db, l)
	e := echo.New()
	bindRouteMiddlewares(e, l)

	// routes
	e.GET("/api/status", h.GetStatus)

	bindRoutes(e, h)
	exportRoutesJson(e)
	return e
}

func bindRoutes(e *echo.Echo, h *handlers.Handlers) {
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
	slogger := logger.Sugar()
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		if len(reqBody) == 0 {
			slogger.Debug("request body: ", "None")
		} else {
			slogger.Debug("request body: ", string(reqBody))
		}

		if len(resBody) == 0 {
			slogger.Debug("response body: ", "No Content")
		} else {
			slogger.Debug("response body: ", string(resBody))
		}
	}))
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
