package userscontroller

import (
	"fmt"
	usersdal "gin-microservice/dal/usersDal"
	usersmodels "gin-microservice/models/usersModels"
	"gin-microservice/utils"
)

type IUsersController interface {
	CreateUser(usersmodels.User) (usersmodels.UserObject, error)
}

type UsersController struct {
	Dal usersdal.IUsersDal
}

func (uc *UsersController) CreateUser(newUser usersmodels.User) (usersmodels.UserObject, error) {
	userPassword, err := utils.ConvertStringToHash(newUser.Password)

	if err != nil {
		return usersmodels.UserObject{}, utils.ErrUnableToHashPasswrod
	}

	newUser.Password = userPassword
	newUser.Username = utils.RemoveSpacesFromInput(newUser.Username)

	createdUser, err := uc.Dal.CreateUser(newUser)

	fmt.Println(err)

	if err != nil {
		return usersmodels.UserObject{}, utils.ErrUnableToRegisterUser
	}

	return createdUser, nil
}
