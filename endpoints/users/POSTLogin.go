package users

import (
	usersmodels "gin-microservice/models/usersModels"
	"gin-microservice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary		Login
//	@Description	Login with Username and Password
//	@Produce		json
//	@Param			body	body		usersmodels.UserLogin	true	"User Login Data"
//	@Success		200		{object}	usersmodels.Token
//	@Failure		400		{string}	string	"invalid user object； invalid to create user"
//	@Router			/users/login [post]
func (ue *UsersEndpoints) Login(c *gin.Context) {
	var userLoginData usersmodels.UserLogin

	if err := c.ShouldBindJSON(&userLoginData); err != nil {
		c.String(http.StatusBadRequest, utils.ErrInvalidUserData.Error())
		return
	}

	var token usersmodels.Token
	var err error

	token.Token, err = ue.Controller.VerifyLogin(userLoginData)

	if err != nil {
		c.String(http.StatusBadRequest, utils.StringWrapError(utils.ErrInvalidLogin, err))
		return
	}

	c.JSON(http.StatusOK, token)
}
