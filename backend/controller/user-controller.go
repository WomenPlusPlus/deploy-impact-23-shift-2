package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context) entity.User
}

type userController struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (controller *userController) FindAll() []entity.User {
	return controller.service.FindAll()
}

func (controller *userController) Save(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON((&user))
	controller.service.Save(user)
	return user
}
