package wishlistservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
)

func (s wishListService) CleanWishList(ctx context.Context, userId uint) (models.Response[bool], error) {
	response := models.Response[bool]{}

	res, err := s.wishListRepo.CleanWishList(ctx, userId)
	if err != nil {
		response.Message = "could not clean wish list"
		return response, err
	}

	response.Data = res
	response.Success = true
	response.Message = "wish list successfully deleted"
	return response, nil
}
