package productservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (s *productService) ListFeaturedProducts(ctx context.Context) (models.Response[[]productmodel.Product], error) {
	response, err := s.productRepo.ListFeaturedProducts(ctx)
	if err != nil {
		response.Success = false
		return response, err
	}
	return response, nil
}
