package productservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (s *productService) GetProduct(ctx context.Context, productId uint) (models.Response[productmodel.Product], error) {
	response := models.Response[productmodel.Product]{}

	if productId == 0 {
		response.Success = false
		return response, ErrProductIdEmpty
	}

	product, err := s.productRepo.GetProduct(ctx, productId)
	if err != nil {
		response.Success = false
		return response, err
	}

	response.Data = product
	response.Success = true
	return response, nil
}
