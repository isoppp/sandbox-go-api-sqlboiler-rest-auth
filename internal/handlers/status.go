package handlers

import (
	"fmt"
	"net/http"
	"sandbox-go-api-sqlboiler-rest-auth/models"

	"github.com/labstack/echo/v4"
)

type CookieValue struct {
	UserID int
	Name   string
}

func (h *Handlers) GetStatus(c echo.Context) error {
	var u *models.User
	uv := c.Get("user")
	if uv != nil {
		u = uv.(*models.User)
		fmt.Println("user data?", u)
	} else {
		fmt.Println("not set user session")
	}
	return c.String(http.StatusOK, "server is running")
}
