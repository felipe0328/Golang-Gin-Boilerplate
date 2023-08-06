package users

import (
	userscontroller "gin-microservice/controller/usersController"
	usersmodels "gin-microservice/models/usersModels"
	"gin-microservice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type POSTLogin struct {
	Controller userscontroller.IUsersController
}

func (pl *POSTLogin) Login (c *gin.Context){
	var userLoginData usersmodels.UserLogin

	if err := c.ShouldBindJSON(&userLoginData); err != nil {
		c.String(http.StatusBadRequest, utils.ErrInvalidUserData.Error())
		return
	}

	token, err :=  pl.Controller.VerifyLogin(userLoginData)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return 
	}

	c.String(http.StatusOK, token)
}