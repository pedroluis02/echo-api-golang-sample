package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	base "github.com/pedroluis02/echo-api-sample1/base"
	repo "github.com/pedroluis02/echo-api-sample1/data"
	model "github.com/pedroluis02/echo-api-sample1/model"
)

func NewAuthService(e *echo.Echo) {
	group := e.Group("/auth")

	group.POST("/token", createToken)
	group.POST("/login", login)

	repo.Init()
}

func createToken(c echo.Context) (err error) {
	credential := new(model.AuthCredential)
	if err = c.Bind(credential); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	auth := &model.Authorization{
		Token: credential.UserName + ":" + credential.Password,
	}
	response := base.CreateResponse(auth, "", nil)

	return c.JSON(http.StatusOK, response)
}

func login(c echo.Context) (err error) {
	credential := new(model.AuthCredential)
	if err = c.Bind(credential); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := repo.FindUserByCredential(credential.UserName, credential.Password)
	response := base.CreateResponse(user, "", err)

	return c.JSON(http.StatusOK, response)
}
