package wishlistrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	wishlistmodel "github.com/amarantec/e-shop/models/wishlist_model"
)

func (r *wishListRepository) AddToWishList(ctx context.Context, item wishlistmodel.WishList) (wishlistmodel.WishList, error) {
	if err :=
		database.DB.WithContext(ctx).
			Create(&item).Error; err != nil {
		return wishlistmodel.WishList{}, err
	}

	return item, nil
}
