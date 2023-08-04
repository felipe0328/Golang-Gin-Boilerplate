package usersmodels

type User struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
