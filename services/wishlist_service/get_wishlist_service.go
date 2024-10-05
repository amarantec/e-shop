package wishlistservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	wishlistmodel "github.com/amarantec/e-shop/models/wishlist_model"
	productrepository "github.com/amarantec/e-shop/repositories/product_repository"
	producttyperepository "github.com/amarantec/e-shop/repositories/product_type_repository"
)

func (s wishListService) GetWishList(ctx context.Context, userId uint) (models.Response[wishlistmodel.WishListResponse], error) {
	response := models.Response[wishlistmodel.WishListResponse]{}

	items, err := s.wishListRepo.GetWishList(ctx, userId)
	if err != nil {
		response.Message = "could not get wish list"
		return response, err
	}

	count, _ := s.wishListRepo.GetWishListCount(ctx, userId)

	wishListResponse := wishlistmodel.WishListResponse{
		QuantityItems: count,
		Products:      []wishlistmodel.WishListProductResponse{},
	}

	for _, item := range items {
		product, _ := productrepository.NewProductRepository().GetProduct(ctx, item.ProductID)
		pt, _ := producttyperepository.NewProductTypeRepository().GetProductType(ctx, item.ProductTypesID)
		wishListProductResponse := wishlistmodel.WishListProductResponse{}
		wishListProductResponse.ProductID = item.ProductID
		wishListProductResponse.Title = product.Title
		wishListProductResponse.ProductType = pt.Name
		wishListProductResponse.ImageURL = product.ImageURL
		wishListResponse.Products = append(wishListResponse.Products, wishListProductResponse)
	}

	response.Data = wishListResponse
	response.Success = true
	response.Message = "wish list successfully retrivied"
	return response, nil
}
