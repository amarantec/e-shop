package userrepository

import (
	"context"

	"github.com/amarantec/e-shop/models"
	usermodel "github.com/amarantec/e-shop/models/user_model"
)

type UserRepository interface {
	Register(ctx context.Context, user *usermodel.User) (models.Response[uint], error)
	FindUser(ctx context.Context, email string) (usermodel.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}
