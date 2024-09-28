package producttypeservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (s *productTypesService) ListProductTypes(ctx context.Context) (models.Response[[]productmodel.ProductTypes], error) {
	response, err := s.productTypesRepo.ListProductTypes(ctx)
	if err != nil {
		response.Success = false
		return response, err
	}

	return response, nil
}
