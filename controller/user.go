package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	base "github.com/pedroluis02/echo-api-sample1/base"
	repo "github.com/pedroluis02/echo-api-sample1/data"
	model "github.com/pedroluis02/echo-api-sample1/model"
)

func NewUserService(e *echo.Echo) {
	group := e.Group("/users")

	group.GET("", getAllUsers)
	group.POST("/create", createUser)
	group.GET("/:id", getUser)
}

func getAllUsers(c echo.Context) (err error) {
	users := repo.GetAllUsers()
	response := base.CreateResponse(users, "User list OK.", nil)

	return c.JSON(http.StatusOK, response)
}

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

func getUser(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := repo.FindUserByIdWithError(uint64(id))
	response := base.CreateResponse(user, "User was found.", err)

	return c.JSON(http.StatusOK, response)
}
