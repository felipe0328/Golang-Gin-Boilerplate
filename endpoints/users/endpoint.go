package users

import (
	usersmodels "gin-microservice/models/usersModels"
	"gin-microservice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context){
	var newUser usersmodels.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.String(http.StatusBadRequest, utils.InvalidUserObject.Error())
		return
	}

	c.JSON(http.StatusOK, newUser)
}