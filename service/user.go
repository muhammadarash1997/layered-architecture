package service

import (
	"layered-architecture/model"
	"layered-architecture/repository"
)

type UserService interface {
	Create(model.User) (model.User, error)
	Update(model.User) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func (this *userService) Create(user model.User) (model.User, error)

func (this *userService) Update(user model.User) (model.User, error)