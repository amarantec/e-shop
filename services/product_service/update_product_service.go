package productservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (s *productService) UpdateProduct(ctx context.Context, product productmodel.Product) (models.Response[bool], error) {
	response := models.Response[bool]{}

	if valid, err := validProduct(product); err != nil || !valid {
		response.Success = false
		return response, ErrProductInvalidFormat
	}

	response, err := s.productRepo.UpdateProduct(ctx, product)
	if err != nil {
		response.Success = false
		return response, err
	}

	return response, nil
}
