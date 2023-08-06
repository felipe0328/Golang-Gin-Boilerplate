package users

import (
	usersmodels "gin-microservice/models/usersModels"
	"gin-microservice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ue *UsersEndpoints) createUser(c *gin.Context) {
	var newUser usersmodels.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.String(http.StatusBadRequest, utils.ErrInvalidUserObject.Error())
		return
	}

	createdUser, err := ue.Controller.CreateUser(newUser)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
