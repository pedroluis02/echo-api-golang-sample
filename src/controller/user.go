package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	base "github.com/pedroluis02/echo-api-sample1/src/base"
	repo "github.com/pedroluis02/echo-api-sample1/src/data"
	model "github.com/pedroluis02/echo-api-sample1/src/model"
)

func NewUserService(e *echo.Echo) {
	group := e.Group("/users")

	group.GET("", getAllUsers)
	group.POST("/create", createUser)
	group.GET("/:id", getUser)
}

// ShowUsers godoc
//	@Summary		Get users
//	@Description	Get all users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200		{array}		model.User
//	@Router			/users	[get]
func getAllUsers(c echo.Context) (err error) {
	users := repo.GetAllUsers()
	response := base.CreateResponse(users, "User list OK.", nil)

	return c.JSON(http.StatusOK, response)
}

// CreateUser godoc
//	@Summary		Create user
//	@Description	Create new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user body model.User true "user info"
//	@Success		200	{object}	model.User
//	@Router			/users/create	[post]
func createUser(c echo.Context) (err error) {
	user := new(model.User)
	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var response *base.BaseObjectResponse

	_, err = repo.CheckIfUserIsUnique(user.Email)
	if err != nil {
		response = base.CreateResponseWithError(err)
	} else {
		newUser, err := repo.AddNewUser(user)
		response = base.CreateResponse(newUser, "User was created.", err)
	}

	return c.JSON(http.StatusOK, response)
}

// ShowUser godoc
//	@Summary		Get user
//	@Description	Get one user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id path integer true "user seach by id"
//	@Success		200	{object}	model.User
//	@Router			/users/{id}	[get]
func getUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := repo.FindUserByIdWithError(uint64(id))
	response := base.CreateResponse(user, "User was found.", err)

	return c.JSON(http.StatusOK, response)
}
