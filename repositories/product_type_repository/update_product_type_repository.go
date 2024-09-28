package producttyperepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (r *productTypesRepository) UpdateProductType(ctx context.Context, productType productmodel.ProductTypes) (models.Response[bool], error) {
	response := models.Response[bool]{}

	result :=
		database.DB.WithContext(ctx).
			Where("id = ?", productType.ID).
			Updates(productmodel.ProductTypes{
				Name: productType.Name,
			})

	if result.Error != nil {
		response.Success = false
		response.Message = "error when updating product type register"
		return response, result.Error
	}

	if result.RowsAffected == 0 {
		response.Success = true
		response.Message = "product type not found"
		return response, nil
	}

	response.Data = true
	response.Success = true
	response.Message = "Product type updated successfully"
	return response, nil
}
