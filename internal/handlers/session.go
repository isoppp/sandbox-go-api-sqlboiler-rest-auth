package handlers

import (
	"net/http"
	"sandbox-go-api-sqlboiler-rest-auth/internal/boilmodels"
	"sandbox-go-api-sqlboiler-rest-auth/internal/cookie"
	"sandbox-go-api-sqlboiler-rest-auth/internal/middleware"
	"strconv"
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

func CreateSession(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	ctx := cc.Request().Context()
	req := new(CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ok, _ := boilmodels.Users(qm.Where("email = ?", req.Email)).Exists(ctx, cc.DB)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	u, err := boilmodels.Users(qm.Where("email = ?", req.Email)).One(ctx, cc.DB)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cc.ZapLogger.Info(u)
	if u.HashedPassword != req.Password {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	expirationDays, err := strconv.Atoi(cc.Config.SessionExpirationDays)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	var s boilmodels.Session
	var uid, _ = uuid.NewUUID()
	s.ID = uid.String()
	s.UserID = u.ID
	s.ExpiresAt = time.Now().Add(time.Hour * 24 * time.Duration(expirationDays))
	err = s.Insert(ctx, cc.DB, boil.Infer())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	encoded, err := cc.SecureCookie.Encode(cookie.SecureCookieSessionKeyName, s.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	ck := &http.Cookie{
		Name:     cookie.SessionCookieKeyName,
		Value:    encoded,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * time.Duration(expirationDays)),
		Secure:   !cc.Config.IsDev,
		HttpOnly: true,
		SameSite: 2,
	}
	c.SetCookie(ck)
	return c.NoContent(http.StatusNoContent)
}

func DeleteSession(c echo.Context) error {
	//cc := c.(*customcontext.CustomContext)
	//ctx := cc.Request().Context()
	return c.NoContent(http.StatusNoContent)
}
