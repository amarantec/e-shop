package userrepository

import (
	"context"

	"github.com/amarantec/e-shop/models"
	usermodel "github.com/amarantec/e-shop/models/user_model"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *usermodel.User) (models.Response[uint], error)
	FindUserByEmail(ctx context.Context, email string) (usermodel.User, error)
	FindUserById(ctx context.Context, userId uint) (usermodel.User, error)
	ChangePassword(ctx context.Context, userId uint, password string) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}
