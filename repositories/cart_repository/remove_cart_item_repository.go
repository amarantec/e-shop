package cartrepository

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
	"gorm.io/gorm"
)

func (r *cartItemsRepository) RemoveItemFromCart(ctx context.Context, productID, productTypesID uint) (models.Response[bool], error) {
	response := models.Response[bool]{}

	if err :=
		database.DB.WithContext(ctx).
			Where("id = ? AND product_types_id = ?", productID, productTypesID).
			Unscoped().Delete(&cartmodel.CartItems{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Success = true
			response.Message = "cart item not found"
			return response, nil
		}
		return response, err
	}

	response.Data = true
	response.Success = true
	response.Message = "Cart Item successfully removed"
	return response, nil
}
