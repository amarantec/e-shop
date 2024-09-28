package producttyperepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (r *productTypesRepository) AddProductType(ctx context.Context, productType productmodel.ProductTypes) (models.Response[productmodel.ProductTypes], error) {
	response := models.Response[productmodel.ProductTypes]{}

	if err :=
		database.DB.WithContext(ctx).
			Create(&productType).Error; err != nil {
		response.Success = false
		return response, err
	}

	response.Data = productType
	response.Success = true
	response.Message = "Product Type successfully registered"
	return response, nil
}
