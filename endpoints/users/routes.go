package users

import (
	"database/sql"
	userscontroller "gin-microservice/controller/usersController"
	usersdal "gin-microservice/dal/usersDal"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, db *sql.DB) {
	usersDal := usersdal.UsersDal{DB: db}
	usersController := &userscontroller.UsersController{Dal: &usersDal}

	postUser := PostUsers{
		Controller: usersController,
	}

	userGroup := r.Group("/users")
	userGroup.POST("/", postUser.createUser)
}
