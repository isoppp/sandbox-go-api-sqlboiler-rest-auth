package handlers

import (
	"errors"
	"net/http"
	"sandbox-go-api-sqlboiler-rest-auth/internal/boilmodels"
	"sandbox-go-api-sqlboiler-rest-auth/internal/middleware"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/labstack/echo/v4"
)

type PublicUser struct {
	ID        int       `boil:"id" json:"id"`
	Email     string    `boil:"email" json:"email"`
	CreatedAt time.Time `boil:"created_at" json:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at"`
}

type UsersData struct {
	Users *[]PublicUser `json:"users"`
}

func GetUsers(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	ctx := cc.Request().Context()
	var users []PublicUser
	res := SuccessResponse{
		Data: UsersData{Users: &users},
	}
	err := boilmodels.Users().Bind(ctx, cc.DB, &users)
	if err != nil {
		c.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, res)
}

type UserData struct {
	User *PublicUser `json:"user"`
}

func GetUser(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	ctx := cc.Request().Context()
	res := SuccessResponse{
		Data: UserData{},
	}
	var id int
	err := echo.PathParamsBinder(c).
		Int("id", &id).
		BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	exists, err := boilmodels.UserExists(ctx, cc.DB, id)
	if !exists {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var users []*PublicUser
	err = boilmodels.Users(qm.Where("id = ?", id)).Bind(ctx, cc.DB, &users)
	if err != nil {
		c.Error(err)
		return err
	}
	res.Data = &users[0]

	return c.JSON(http.StatusOK, res)
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	ctx := cc.Request().Context()
	res := SuccessResponse{
		Data: UserData{},
	}
	req := new(CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	u := boilmodels.User{
		Email:          req.Email,
		HashedPassword: req.Password,
	}
	err := u.Insert(ctx, cc.DB, boil.Infer())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	pu := PublicUser{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	res.Data = UserData{
		User: &pu,
	}

	return c.JSON(http.StatusOK, res)
}

func PatchUser(c echo.Context) error {
	return errors.New("not implemented")
}

func DeleteUser(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	ctx := cc.Request().Context()
	var id int
	err := echo.PathParamsBinder(c).
		Int("id", &id).
		BindError()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	exists, err := boilmodels.UserExists(ctx, cc.DB, id)
	if !exists {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	u, err := boilmodels.FindUser(ctx, cc.DB, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	_, err = u.Delete(ctx, cc.DB)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
