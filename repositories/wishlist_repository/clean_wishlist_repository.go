package wishlistrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	wishlistmodel "github.com/amarantec/e-shop/models/wishlist_model"
)

func (r *wishListRepository) CleanWishList(ctx context.Context, userId uint) (bool, error) {
	if err :=
		database.DB.WithContext(ctx).
			Where("user_id = ?", userId).
			Unscoped().Delete(&wishlistmodel.WishList{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
