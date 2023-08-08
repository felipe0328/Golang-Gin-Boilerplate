package users

import (
	"database/sql"
	userscontroller "gin-microservice/controller/usersController"
	usersdal "gin-microservice/dal/usersDal"
	"gin-microservice/endpoints/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersEndpoints struct {
	Controller userscontroller.IUsersController
}

func UserRoutes(r *gin.Engine, db *sql.DB) {
	usersDal := usersdal.UsersDal{DB: db}
	usersController := &userscontroller.UsersController{Dal: &usersDal}

	endpoints := UsersEndpoints{Controller: usersController}

	userGroup := r.Group("/users")
	userGroup.POST("/", endpoints.CreateUser)
	userGroup.POST("/login", endpoints.Login)

	authenticated := r.Group("/users")
	authenticated.Use(middlewares.JwtAuthMiddleare)
	authenticated.GET("/ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })
}
