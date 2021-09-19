package handlers

import (
	"context"
	"net/http"
	"sandbox-go-api-sqlboiler-rest-auth/internal/scookie"
	"sandbox-go-api-sqlboiler-rest-auth/models"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/labstack/echo/v4"
)

type CreateSessionRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handlers) CreateSession(c echo.Context) error {
	ctx := context.Background()
	req := new(CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ok, _ := models.Users(qm.Where("email = ?", req.Email)).Exists(ctx, h.db)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	u, err := models.Users(qm.Where("email = ?", req.Email)).One(ctx, h.db)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	h.slogger.Info(u)
	if u.HashedPassword != req.Password {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	var s models.Session
	var uid, _ = uuid.NewUUID()
	s.ID = uid.String()
	s.UserID = u.ID
	s.ExpiresAt = time.Now().Add(time.Hour * 24 * 30)
	err = s.Insert(ctx, h.db, boil.Infer())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	sc := scookie.NewSecureCookie()
	encoded, err := sc.Encode("session", s.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cookie := &http.Cookie{
		Name:     "session",
		Value:    encoded,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		Secure:   true,
		HttpOnly: true,
		SameSite: 2,
	}
	c.SetCookie(cookie)
	return c.NoContent(http.StatusNoContent)
}

func (h *Handlers) DeleteSession(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
