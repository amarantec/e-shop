package cartrepository

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/config/database"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
	"gorm.io/gorm"
)

func (r *cartItemsRepository) GetCartItem(ctx context.Context, cartItem cartmodel.CartItems) (cartmodel.CartItems, error) {
	cartRetrivied := cartmodel.CartItems{}
	if err :=
		database.DB.WithContext(ctx).
			Where("user_id = ? AND product_id = ? AND product_types_id = ?", cartItem.UserID, cartItem.ProductID, cartItem.ProductTypesID).
			First(&cartRetrivied).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cartmodel.CartItems{}, nil
		}
		return cartmodel.CartItems{}, err
	}
	return cartRetrivied, nil
}
