package userscontroller

import (
	usersmodels "gin-microservice/models/usersModels"
	"gin-microservice/utils"
)

func (uc *UsersController) VerifyLogin(userLoginData usersmodels.UserLogin) (string, error) {

	userDalObject, err := uc.Dal.GetUserByLoginData(userLoginData)

	if err != nil {
		return "", err
	}

	err = utils.VerifyPassword(userLoginData.Password, userDalObject.Password)

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(userDalObject.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
