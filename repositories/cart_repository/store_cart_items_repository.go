package cartrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
)

func (r *cartItemsRepository) StoreCartItems(ctx context.Context, cartItems cartmodel.CartItems) (models.Response[cartmodel.CartItems], error) {
	response := models.Response[cartmodel.CartItems]{}
	if err :=
		database.DB.WithContext(ctx).
			Create(&cartItems).Error; err != nil {
		return response, err
	}

	response.Data = cartItems
	response.Success = true
	response.Message = "Item successfully stored"

	return response, nil
}
