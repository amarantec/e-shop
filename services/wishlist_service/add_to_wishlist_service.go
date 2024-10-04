package wishlistservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	wishlistmodel "github.com/amarantec/e-shop/models/wishlist_model"
)

func (s wishListService) AddToWishList(ctx context.Context, item wishlistmodel.WishList) (models.Response[wishlistmodel.WishList], error) {
	response := models.Response[wishlistmodel.WishList]{}

	res, err := s.wishListRepo.AddToWishList(ctx, item)
	if err != nil {
		response.Message = "could not add this item to wish list"
		return response, err
	}

	response.Data = res
	response.Success = true
	response.Message = "item successfully added to wish list"
	return response, nil
}
