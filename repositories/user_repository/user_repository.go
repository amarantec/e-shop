package userrepository

import (
	usermodel "github.com/amarantec/project777/models/user_model"
)

type UserRepository interface {
	Register(user *usermodel.User) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}
