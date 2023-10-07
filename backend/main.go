package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/middlewares"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

var (
	userService    service.UserService       = service.New()
	userController controller.UserController = controller.New(userService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})

	server.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.FindAll())
	})

	server.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(200, userController.Save(ctx))
	})

	// Start server
	server.Run(":8080")
}
