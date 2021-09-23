package handlers

import (
	"net/http"
	"sandbox-go-api-sqlboiler-rest-auth/internal/middleware"

	"github.com/labstack/echo/v4"
)

func GetStatus(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	u := cc.CurrentUser
	if u != nil {
		cc.ZapLogger.Debug("user data?", u)
	} else {
		cc.ZapLogger.Debug("not set user session")
	}
	return c.String(http.StatusOK, "server is running")
}
