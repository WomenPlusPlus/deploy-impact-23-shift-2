package service

import "gitlab.com/pragmaticreviews/golang-gin-poc/entity"

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
}

type userService struct {
	users []entity.User
}

func New() UserService {
	return &userService{
		users: []entity.User{},
	}
}

func (service *userService) Save(user entity.User) entity.User {
	service.users = append(service.users, user)
	return user
}

func (service *userService) FindAll() []entity.User {
	return service.users
}
