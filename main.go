package main

import (
	"database/sql"
	"gin-microservice/dal"
	"gin-microservice/endpoints/users"
	dalmodels "gin-microservice/models/dalModels"
	"gin-microservice/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	docs "gin-microservice/docs"
)

const (
	DB_HOST = "DB_HOST"
	DB_PORT = "DB_PORT"
	DB_USER = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_DATABASE_NAME = "DB_DATABASE_NAME"
	DB_DATABASE_DRIVER = "DB_DATABASE_DRIVER"
)

func main() {
	godotenv.Load(".env")

	r := gin.Default()

	db := connectToDB()
	registerEndpoints(r, db)

	err := r.Run()

	utils.CheckPanic(err)
}

func connectToDB()*sql.DB {
	dbConnection := dalmodels.DBConnection{
		Host: os.Getenv(DB_HOST),
		Port: os.Getenv(DB_PORT),
		User: os.Getenv(DB_USER),
		Password: os.Getenv(DB_PASSWORD),
		DatabaseName: os.Getenv(DB_DATABASE_NAME),
		DatabaseDriver: os.Getenv(DB_DATABASE_DRIVER),
	}

	return dal.GetDB(dbConnection)
}

func registerEndpoints(r *gin.Engine, db *sql.DB){
	registerPing(r)
	showSwaggerDocumentation(r)
	users.UserRoutes(r, db)
}

func registerPing(r *gin.Engine){
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func showSwaggerDocumentation(r *gin.Engine) {
	docs.SwaggerInfo.Description = "Microservice Documentation"

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}