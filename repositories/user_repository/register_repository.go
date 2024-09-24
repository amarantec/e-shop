package userrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	usermodel "github.com/amarantec/e-shop/models/user_model"
)

func (r *userRepository) SaveUser(ctx context.Context, user *usermodel.User) (models.Response[uint], error) {
	response := models.Response[uint]{}

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
