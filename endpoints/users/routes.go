package users

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.Engine){
	userGroup := r.Group("/user")
	userGroup.POST("/", createUser)
}