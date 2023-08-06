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

	postCreateUser := PostCreateUser{Controller: usersController}
	postLogin := POSTLogin{Controller: usersController}

	userGroup := r.Group("/users")
	userGroup.POST("/", postCreateUser.createUser)
	userGroup.POST("/login", postLogin.Login)
}
