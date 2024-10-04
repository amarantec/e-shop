package wishlistservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	wishlistmodel "github.com/amarantec/e-shop/models/wishlist_model"
	wishlistrepository "github.com/amarantec/e-shop/repositories/wishlist_repository"
)

type WishListService interface {
	AddToWishList(ctx context.Context, item wishlistmodel.WishList) (models.Response[wishlistmodel.WishList], error)
	CleanWishList(ctx context.Context, userId uint) (models.Response[bool], error)
	GetWishList(ctx context.Context, userId uint) (models.Response[wishlistmodel.WishListResponse], error)
	RemoveFromWishList(ctx context.Context, userId, itemId uint) (models.Response[bool], error)
}

type wishListService struct {
	wishListRepo wishlistrepository.WishListRepository
}

func NewWishListService(repo wishlistrepository.WishListRepository) WishListService {
	return &wishListService{wishListRepo: repo}
}
