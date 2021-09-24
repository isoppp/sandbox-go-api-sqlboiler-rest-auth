package handlers

import (
	"errors"
	"net/http"
	"sandbox-go-api-sqlboiler-rest-auth/internal/boilmodels"
	"sandbox-go-api-sqlboiler-rest-auth/internal/middleware"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/labstack/echo/v4"
)

type PublicUser struct {
	ID    int      `json:"id"`
	Email string   `json:"email"`
	Roles []string `json:"roles,omitempty"`
}

type UsersData struct {
	Users *[]PublicUser `json:"users"`
}

type UserData struct {
	User *PublicUser `json:"user"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Me(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	cu := cc.CurrentUser
	cc.ZapLogger.Info("cu:", cu)
	if cu == nil {
		return echo.NewHTTPError(http.StatusForbidden, http.StatusText(http.StatusForbidden))
	}

	var roles []string
	for i := range cu.R.Roles {
		roles = append(roles, cu.R.Roles[i].Name)
	}

	return c.JSON(http.StatusOK, JsonSuccessResponse(UserData{
		User: &PublicUser{
			ID:    cu.ID,
			Email: cu.Email,
			Roles: roles,
		},
	}))
}

func GetUsers(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	ctx := cc.Request().Context()
	var users []PublicUser
	err := boilmodels.Users().Bind(ctx, cc.DB, &users)
	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, JsonSuccessResponse(UsersData{
		Users: &users,
	}))
}

func GetUser(c echo.Context) error {
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

	var users []*PublicUser
	err = boilmodels.Users(qm.Where("id = ?", id)).Bind(ctx, cc.DB, &users)
	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, JsonSuccessResponse(UserData{
		User: users[0],
	}))
}

func CreateUser(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	ctx := cc.Request().Context()
	req := new(CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	r, err := boilmodels.Roles(qm.Where("name = ?", boilmodels.UserRoleTypeUser)).One(ctx, cc.DB)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	cc.ZapLogger.Info(r.Name)

	u := boilmodels.User{
		Email:          req.Email,
		HashedPassword: req.Password,
	}

	err = u.Insert(ctx, cc.DB, boil.Infer())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = u.SetRoles(ctx, cc.DB, false, r)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	var roles []string
	for i := range u.R.Roles {
		roles = append(roles, u.R.Roles[i].Name)
	}
	pu := PublicUser{
		ID:    u.ID,
		Email: u.Email,
		Roles: roles,
	}

	return c.JSON(http.StatusOK, JsonSuccessResponse(UserData{
		User: &pu,
	}))
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
