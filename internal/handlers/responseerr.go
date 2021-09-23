package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type errorContent struct {
	Message interface{} `json:"message"`
}
type errorResponse struct {
	Error interface{} `json:"error"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(he.Code, errorResponse{
				Error: errorContent{
					Message: he.Message,
				},
			})
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}
}
