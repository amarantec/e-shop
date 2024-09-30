package cartrepository

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
	"gorm.io/gorm"
)

func (r *cartItemsRepository) UpdateQuantity(ctx context.Context, cartItem cartmodel.CartItems) (models.Response[bool], error) {
	response := models.Response[bool]{}

	if err :=
		database.DB.WithContext(ctx).Model(cartmodel.CartItems{}).
			Where("id = ? AND user_id = ? AND product_types_id = ?", cartItem.ID, cartItem.UserID, cartItem.ProductTypesID).
			Update("quantity", gorm.Expr("quantity + ?", cartItem.Quantity)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Message = "Cart item not found"
			return response, nil
		}
		return response, err
	}

	response.Data = true
	response.Success = true
	response.Message = "Cart item quantity successfully updated"
	return response, nil
}
