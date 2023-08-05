package users

import (
	userscontroller "gin-microservice/controller/usersController"
	usersmodels "gin-microservice/models/usersModels"
	"gin-microservice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostUsers struct {
	Controller userscontroller.IUsersController
}

func (ue *PostUsers) createUser(c *gin.Context) {
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

	c.JSON(http.StatusOK, createdUser)
}
