package usersdal

import usersmodels "gin-microservice/models/usersModels"

func (ud *UsersDal) CreateUser(newUser usersmodels.User)(usersmodels.UserObject, error){
	query := `INSERT into USERS(username, password, email, first_name, last_name) VALUES ($1, $2, $3, $4, $5) returning id, username, first_name , last_name , email, created_on , is_active`

	result, err := ud.Query(
		query,newUser.Username, 
		newUser.Password, 
		newUser.Email, 
		newUser.FirstName, 
		newUser.LastName,
	)

	if err != nil {
		return usersmodels.UserObject{}, err
	}

	defer result.Close()

	createdObject := usersmodels.UserObject{}

	for result.Next() {
		err := result.Scan(
			&createdObject.ID,
			&createdObject.Username,
			&createdObject.FirstName,
			&createdObject.LastName,
			&createdObject.Email,
			&createdObject.CreatedOn,
			&createdObject.IsActive,
		)

		if err != nil {
			return usersmodels.UserObject{}, err
		}
	}

	return createdObject, nil
}