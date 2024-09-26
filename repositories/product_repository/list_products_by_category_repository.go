package productrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (r *productRepository) ListProductsByCategory(ctx context.Context, categoryUrl string) (models.Response[[]productmodel.Product], error) {
	response := models.Response[[]productmodel.Product]{}

	result :=
		database.DB.WithContext(ctx).
			Model(&productmodel.Product{}).
			Preload("Category").
			Preload("ProductVariants", "visible = ? AND deleted = ?", true, false).
			Preload("Images").
			Where("url = ?", categoryUrl).Find(&response.Data)

	if result.Error != nil {
		response.Success = false
		return response, result.Error
	}

	response.Success = true
	response.Message = "Products retrivied successfully"
	return response, nil
}
