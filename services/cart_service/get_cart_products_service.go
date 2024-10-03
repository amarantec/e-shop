package cartservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
	productrepository "github.com/amarantec/e-shop/repositories/product_repository"
)

func (s *cartService) GetCartProducts(ctx context.Context, userId uint) (models.Response[[]cartmodel.CartProductResponse], error) {
	response := models.Response[[]cartmodel.CartProductResponse]{}

	cartItems, err := s.cartRepo.GetCartProducts(ctx, userId)
	if err != nil {
		return response, err
	}

	for _, item := range cartItems {
		product, _ := productrepository.NewProductRepository().GetProduct(ctx, item.ProductID)
		for _, i := range product.ProductVariants {
			cp := cartmodel.CartProductResponse{
				ProductID:      product.ID,
				Title:          product.Title,
				ImageUrl:       product.ImageURL,
				Price:          i.Price,
				ProductTypes:   i.ProductTypes.Name,
				ProductTypesID: i.ProductTypesID,
				Quantity:       item.Quantity,
			}
			response.Data = append(response.Data, cp)
		}
	}
	response.Success = true
	return response, nil

}
