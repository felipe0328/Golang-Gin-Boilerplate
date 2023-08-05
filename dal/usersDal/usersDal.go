package usersdal

import (
	"database/sql"
	usersmodels "gin-microservice/models/usersModels"
)

type IUsersDal interface {
	CreateUser(usersmodels.User)(usersmodels.UserObject,error)
}

type UsersDal struct {
	*sql.DB
}