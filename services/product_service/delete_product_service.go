package productservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
)

func (s *productService) DeleteProduct(ctx context.Context, productId uint) (models.Response[bool], error) {
	response := models.Response[bool]{}

	if productId == 0 {
		response.Success = false
		return response, ErrProductIdEmpty
	}

	response, err := s.productRepo.DeleteProduct(ctx, productId)
	if err != nil {
		response.Success = false
		return response, err
	}

	return response, nil
}
