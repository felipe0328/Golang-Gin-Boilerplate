package userscontroller

import (
	usersmodels "gin-microservice/models/usersModels"
	"gin-microservice/utils"
)

func (uc *UsersController) VerifyLogin(userLoginData usersmodels.UserLogin) (string, error) {

	userDalObject, err := uc.Dal.GetUserByLoginData(userLoginData)

	if err != nil {
		utils.LogError(utils.WrapError(utils.ErrUserNotFound, err))
		return "", utils.ErrUserNotFound
	}

	err = utils.VerifyPassword(userLoginData.Password, userDalObject.Password)

	if err != nil {
		utils.LogError(utils.WrapError(utils.ErrInvalidPassword, err))
		return "", utils.ErrInvalidPassword
	}

	token, err := utils.GenerateToken(userDalObject.ID)

	if err != nil {
		utils.LogError(utils.WrapError(utils.ErrUnableToGenerateToken, err))
		return "", utils.ErrUnableToGenerateToken
	}

	return token, nil
}
