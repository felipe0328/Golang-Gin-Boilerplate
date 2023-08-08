package users

import (
	usersmodels "gin-microservice/models/usersModels"
	"gin-microservice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ue *UsersEndpoints) CreateUser(c *gin.Context) {
	var newUser usersmodels.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.String(http.StatusBadRequest, utils.ErrInvalidUserObject.Error())
		return
	}

	createdUser, err := ue.Controller.CreateUser(newUser)

	if err != nil {
		c.String(http.StatusBadRequest, utils.StringWrapError(utils.ErrUnableToCreateUser, err))
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
