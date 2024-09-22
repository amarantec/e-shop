package userservice

import (
	usermodel "github.com/amarantec/project777/models/user_model"
	userrepository "github.com/amarantec/project777/repositories/user_repository"
)

type UserService interface {
	Register(user *usermodel.User) error
}

type userService struct {
	userRepo userrepository.UserRepository
}

func NewUserService(repo userrepository.UserRepository) UserService {
	return &userService{userRepo: repo}
}
