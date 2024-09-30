package cartrepository

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
	"gorm.io/gorm"
)

func (r *cartItemsRepository) GetCartItemsCount(ctx context.Context, userId uint) (models.Response[int64], error) {
	response := models.Response[int64]{}

	if err :=
		database.DB.WithContext(ctx).
			Model(cartmodel.CartItems{}).
			Where("user_id = ?", userId).Count(&response.Data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Success = true
			response.Message = "No cart items related to this user id"
			return response, nil
		}
		response.Success = false
		return response, err
	}

	response.Success = true
	response.Message = "Cart items count retrivied successfully"
	return response, nil
}
