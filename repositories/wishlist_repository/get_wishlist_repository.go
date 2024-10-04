package wishlistrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	wishlistmodel "github.com/amarantec/e-shop/models/wishlist_model"
)

func (r wishListRepository) GetWishList(ctx context.Context, userId uint) ([]wishlistmodel.WishList, error) {
	wishList := []wishlistmodel.WishList{}

	if err :=
		database.DB.WithContext(ctx).
			Where("user_id = ?", userId).
			Find(&wishList).Error; err != nil {
		return []wishlistmodel.WishList{}, err
	}
	return wishList, nil
}
