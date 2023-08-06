package users

import (
	usersmodels "gin-microservice/models/usersModels"
	"gin-microservice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ue *UsersEndpoints) Login (c *gin.Context){
	var userLoginData usersmodels.UserLogin

	if err := c.ShouldBindJSON(&userLoginData); err != nil {
		c.String(http.StatusBadRequest, utils.ErrInvalidUserData.Error())
		return
	}

	token, err :=  ue.Controller.VerifyLogin(userLoginData)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return 
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}