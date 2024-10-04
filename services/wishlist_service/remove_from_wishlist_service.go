package wishlistservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
)

func (s wishListService) RemoveFromWishList(ctx context.Context, userId, itemId uint) (models.Response[bool], error) {
	response := models.Response[bool]{}

	res, err := s.wishListRepo.RemoveFromWishList(ctx, userId, itemId)
	if err != nil {
		response.Message = "could not remove this item from wish list"
		return response, err
	}

	response.Data = res
	response.Success = true
	response.Message = "item successfully removed"
	return response, nil
}
