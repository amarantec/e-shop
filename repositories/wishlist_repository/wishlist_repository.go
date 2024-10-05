package wishlistrepository

import (
	"context"

	wishlistmodel "github.com/amarantec/e-shop/models/wishlist_model"
)

type WishListRepository interface {
	AddToWishList(ctx context.Context, item wishlistmodel.WishList) (wishlistmodel.WishList, error)
	CleanWishList(ctx context.Context, userId uint) (bool, error)
	GetWishList(ctx context.Context, userId uint) ([]wishlistmodel.WishList, error)
	RemoveFromWishList(ctx context.Context, userId, itemId uint) (bool, error)
	GetWishListCount(ctx context.Context, userId uint) (int64, error)
}

type wishListRepository struct{}

func NewWishListRepository() WishListRepository {
	return &wishListRepository{}
}
