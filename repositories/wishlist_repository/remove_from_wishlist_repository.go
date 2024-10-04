package wishlistrepository

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/config/database"
	wishlistmodel "github.com/amarantec/e-shop/models/wishlist_model"
	"gorm.io/gorm"
)

func (r *wishListRepository) RemoveFromWishList(ctx context.Context, userId, itemId uint) (bool, error) {
	if err :=
		database.DB.WithContext(ctx).
			Where("user_id = ? AND id = ?", userId, itemId).
			Unscoped().Delete(&wishlistmodel.WishList{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
