package productservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (s *productService) ListProductsByCategory(ctx context.Context, categoryUrl string) (models.Response[[]productmodel.Product], error) {
	response := models.Response[[]productmodel.Product]{}

	if categoryUrl == "" {
		response.Success = false
		return response, ErrProductCategoryUrlEmpty


	}
	
	response, err := s.productRepo.ListProductsByCategory(ctx, categoryUrl)
	if err != nil {
		response.Success = false
		return response, err
	}
	return response, nil
}
