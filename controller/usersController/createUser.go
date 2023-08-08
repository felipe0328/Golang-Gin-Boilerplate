package userscontroller

import (
	usersmodels "gin-microservice/models/usersModels"
	"gin-microservice/utils"
)

func (uc *UsersController) CreateUser(newUser usersmodels.User) (usersmodels.UserObject, error) {
	userPassword, err := utils.ConvertStringToHash(newUser.Password)

	if err != nil {
		utils.LogError(utils.WrapError(utils.ErrUnableToHashPassword, err))
		return usersmodels.UserObject{}, utils.ErrUnableToHashPassword
	}

	newUser.Password = userPassword
	newUser.Username = utils.RemoveSpacesFromInput(newUser.Username)

	createdUser, err := uc.Dal.CreateUser(newUser)

	if err != nil {
		utils.LogError(utils.WrapError(utils.ErrUnableToRegisterUser, err))
		return usersmodels.UserObject{}, utils.ErrUnableToRegisterUser
	}

	return createdUser, nil
}
