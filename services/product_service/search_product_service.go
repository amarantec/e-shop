package productservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (s *productService) SearchProducts(ctx context.Context, searchText string) (models.Response[[]productmodel.Product], error) {
	response := models.Response[[]productmodel.Product]{}

	if searchText == "" {
		response.Success = false
		return response, ErrProductSearchTextEmpty
	}

	response, err := s.productRepo.SearchProducts(ctx, searchText)
	if err != nil {
		response.Success = false
		return response, err
	}
	return response, nil
}
