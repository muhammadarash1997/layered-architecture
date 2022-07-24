package repository

import (
	"layered-architecture/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(model.User) (model.User, error)
	Update(model.User) (model.User, error)
}

type userRepository struct {
	db *mongo.Database
}

func (this *userRepository) Create(user model.User) (model.User, error)

func (this *userRepository) Update(user model.User) (model.User, error)