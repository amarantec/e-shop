package productrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (r *productRepository) ListProducts(ctx context.Context) (models.Response[[]productmodel.Product], error) {
	response := models.Response[[]productmodel.Product]{}
	result :=
		database.DB.WithContext(ctx).
			Model(&productmodel.Product{}).
			Preload("ProductVariants").
			Preload("Images").
			Find(&response.Data)

	if result.Error != nil {
		response.Success = false
		response.Message = "error when get featured products"
		return response, result.Error
	}
	response.Success = true
	response.Message = "Products retrivied successfully"
	return response, nil
}
