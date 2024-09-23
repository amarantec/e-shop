package userrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	usermodel "github.com/amarantec/e-shop/models/user_model"
	"github.com/amarantec/e-shop/utils"
)

func (r *userRepository) Register(ctx context.Context, user *usermodel.User) (models.Response[uint], error) {
	response := models.Response[uint]{}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		response.Success = false
		response.Message = "error when hashing password"
		return response, err
	}

	user.Password = hashedPassword

	if err := database.DB.WithContext(ctx).
		Create(user).Error; err != nil {
		response.Success = false
		return response, err
	}

	response.Data = user.ID
	response.Success = true
	response.Message = "User registred successfully"
	return response, nil
}
