package producttypeservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (s *productTypesService) UpdateProductType(ctx context.Context, productTypes productmodel.ProductTypes) (models.Response[bool], error) {
	response := models.Response[bool]{}
	if productTypes.ID <= 0 {
		response.Success = false
		response.Message = "product type id is empty"
		return response, ErrProductTypesIdEmpty
	}

	if productTypes.Name == "" {
		response.Success = false
		response.Message = "product types name is empty"
		return response, ErrProductTypesNameEmpty
	}

	response, err := s.productTypesRepo.UpdateProductType(ctx, productTypes)
	if err != nil {
		response.Success = false
		return response, err
	}

	return response, nil
}
