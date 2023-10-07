// Package controller provides the implementation of the UserController interface
// for handling user-related operations in the Gin web framework.
package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

// UserController defines the methods that must be implemented by a user controller.
type UserController interface {
	// FindAll retrieves a list of all users.
	FindAll() []entity.User

	// Save saves a user entity based on the request data in the given context.
	Save(ctx *gin.Context) entity.User
}

// userController is the implementation of the UserController interface.
type userController struct {
	service service.UserService
}

// New creates a new instance of UserController with the provided UserService.
func New(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

// FindAll retrieves a list of all users.
func (controller *userController) FindAll() []entity.User {
	return controller.service.FindAll()
}

// Save saves a user entity based on the request data in the given context.
func (controller *userController) Save(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON((&user))
	controller.service.Save(user)
	return user
}
