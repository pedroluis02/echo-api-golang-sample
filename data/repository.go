package data

import (
	"errors"
	"fmt"

	model "github.com/pedroluis02/echo-api-sample1/model"
)

var users []model.User
var emptyUser *model.User

func searchOnList[T any](list []T, compare func(a T) bool) (bool, T) {
	for _, element := range list {
		if compare(element) {
			return true, element
		}
	}

	var empty T
	return false, empty
}

func addUserToList(user *model.User) (bool, *model.User) {
	if user.Id < 1 {
		user.Id = uint64(len(users) + 1)
	} else {
		result, _ := FindUserById(user.Id)
		if result {
			return false, &model.User{}
		}
	}

	users = append(users, *user)
	return true, user
}

func Init() {
	user1 := &model.User{
		Name:     "Admin",
		LastName: "Admin",
		Email:    "admin@example.com",
		Password: "123456",
	}
	addUserToList(user1)
}

func GetAllUsers() []model.User {
	return users
}

func FindUserById(id uint64) (bool, model.User) {
	compare := func(user model.User) bool {
		return id == user.Id
	}

	return searchOnList(users, compare)
}

func FindUserByIdWithError(id uint64) (model.User, error) {
	result, user := FindUserById(id)

	if result {
		return user, nil
	} else {
		return model.User{}, errors.New(fmt.Sprintf("There is no user with id=%d", id))
	}
}

func FindUserByCredential(username string, password string) (*model.User, error) {
	compare := func(user model.User) bool {
		return username == user.Email && password == user.Password
	}

	result, user := searchOnList(users, compare)
	if result {
		return &user, nil
	} else {
		return &model.User{}, errors.New("Username or password is incorrect!")
	}
}

func CheckIfUserIsUnique(username string) (*model.User, error) {
	compare := func(user model.User) bool {
		return username == user.Email
	}

	result, user := searchOnList(users, compare)
	if result {
		return &model.User{}, errors.New(fmt.Sprintf("There is no user with username=%s", username))
	} else {
		return &user, nil
	}
}

func AddNewUser(user *model.User) (*model.User, error) {
	result, newUser := addUserToList(user)
	if result {
		return newUser, nil
	} else {
		return &model.User{}, errors.New("User already exists or data is invalid")
	}
}
