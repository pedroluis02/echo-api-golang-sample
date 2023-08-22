package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	base "github.com/pedroluis02/echo-api-sample1/src/base"
	repo "github.com/pedroluis02/echo-api-sample1/src/data"
	model "github.com/pedroluis02/echo-api-sample1/src/model"
)

func NewAuthService(e *echo.Echo) {
	group := e.Group("/auth")

	group.POST("/token", createToken)
	group.POST("/login", login)

	repo.Init()
}

// CreateToken godoc
//	@Summary		Create token
//	@Description	Create new user token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}		model.Authorization
//	@Param			userAuth body model.AuthCredential true "user info"
//	@Router			/auth/token	[post]
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

// Login godoc
//	@Summary		Login
//	@Description	User login
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}		model.User
//	@Param			userAuth body model.AuthCredential true "user login"
//	@Router			/auth/login	[post]
func login(c echo.Context) (err error) {
	credential := new(model.AuthCredential)
	if err = c.Bind(credential); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := repo.FindUserByCredential(credential.UserName, credential.Password)
	response := base.CreateResponse(user, "", err)

	return c.JSON(http.StatusOK, response)
}
