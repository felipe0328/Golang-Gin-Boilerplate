package usersmodels

type UserLogin struct {
	Username string `json:"username" bind:"required"`
	Password string `json:"password" bind:"required"`
}