package wishlistrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	wishlistmodel "github.com/amarantec/e-shop/models/wishlist_model"
)

func (r *wishListRepository) GetWishListCount(ctx context.Context, userId uint) (int64, error) {
	var count int64

	if err :=
		database.DB.WithContext(ctx).
			Model(wishlistmodel.WishList{}).
			Where("user_id = ?", userId).
			Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
