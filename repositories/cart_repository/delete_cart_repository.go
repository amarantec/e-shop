package cartrepository

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/config/database"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
	"gorm.io/gorm"
)

func (r *cartItemsRepository) DeleteCart(ctx context.Context, userId uint) (bool, error) {
	if err :=
		database.DB.WithContext(ctx).
			Where("user_id = ?", userId).
			Unscoped().Delete(&cartmodel.CartItems{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
