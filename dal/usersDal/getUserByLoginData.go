package usersdal

import (
	usersmodels "gin-microservice/models/usersModels"
)

func (ud *UsersDal) GetUserByLoginData(userLoginData usersmodels.UserLogin) (usersmodels.UserDalObject, error) {
	query := `SELECT * FROM USERS WHERE username = $1`

	var userDalObject usersmodels.UserDalObject

	err := ud.QueryRow(query, userLoginData.Username).Scan(
		&userDalObject.ID,
		&userDalObject.Username,
		&userDalObject.Password,
		&userDalObject.Email,
		&userDalObject.FirstName,
		&userDalObject.LastName,
		&userDalObject.CreatedOn,
		&userDalObject.IsActive,
	)

	return userDalObject, err
}