package producttyperepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (r *productTypesRepository) ListProductTypes(ctx context.Context) (models.Response[[]productmodel.ProductTypes], error) {
	response := models.Response[[]productmodel.ProductTypes]{}

	if err :=
		database.DB.WithContext(ctx).
			Model(productmodel.ProductTypes{}).
			Find(&response.Data).Error; err != nil {
		response.Success = false
		response.Message = "error when getting product types"
		return response, err
	}

	response.Success = true
	response.Message = "Product types retrieved successfully"
	return response, nil
}
