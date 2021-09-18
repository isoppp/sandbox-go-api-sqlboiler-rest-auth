package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) GetStatus(c echo.Context) error {
	return c.String(http.StatusOK, "server is running")
}
