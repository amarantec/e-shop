package cartrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
)

func (r cartItemsRepository) GetCartProducts(ctx context.Context, userId uint) ([]cartmodel.CartItems, error) {
	cart := []cartmodel.CartItems{}

	if err := database.DB.WithContext(ctx).
		Where("user_id = ?", userId).
		Find(&cart).Error; err != nil {
		return []cartmodel.CartItems{}, err
	}

	return cart, nil
}
