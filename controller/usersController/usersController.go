package userscontroller

import (
	usersdal "gin-microservice/dal/usersDal"
	usersmodels "gin-microservice/models/usersModels"
)

type IUsersController interface {
	CreateUser(usersmodels.User) (usersmodels.UserObject, error)
	VerifyLogin(usersmodels.UserLogin)(string, error)
}

type UsersController struct {
	Dal usersdal.IUsersDal
}
