package main

import (
	"gin-microservice/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	err := r.Run()

	utils.CheckPanic(err)
}