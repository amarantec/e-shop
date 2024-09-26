package productrepository

import (
	"context"
	"strings"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (r *productRepository) SearchProducts(ctx context.Context, searchText string) (models.Response[[]productmodel.Product], error) {
	response := models.Response[[]productmodel.Product]{}

	result :=
		database.DB.WithContext(ctx).
			Model(&productmodel.Product{}).
			Preload("ProductVariants", "visible = ? AND deleted = ?", true, false).
			Preload("Images").
			Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ? AND visible = ? AND deleted = ?", strings.ToLower("%"+searchText+"%"),
				strings.ToLower("%"+searchText+"%"), true, false).Find(&response.Data)

	if result.Error != nil {
		response.Success = false
		return response, result.Error
	}
	response.Success = true
	response.Message = "Products retrivied successfully"
	return response, nil
}
