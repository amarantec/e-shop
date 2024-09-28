package producttypeservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (s *productTypesService) AddProductType(ctx context.Context, productTypes productmodel.ProductTypes) (models.Response[productmodel.ProductTypes], error) {
	response := models.Response[productmodel.ProductTypes]{}
	if productTypes.Name == "" {
		response.Success = false
		response.Message = "product types name is empty"
		return response, ErrProductTypesNameEmpty
	}

	response, err := s.productTypesRepo.AddProductType(ctx, productTypes)
	if err != nil {
		response.Success = false
		return response, err
	}

	return response, nil
}
