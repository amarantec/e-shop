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

	if err :=
		database.DB.WithContext(ctx).
			Model(productmodel.Product{}).
			Preload("Category").                     // Precarregar a categoria associada
			Preload("ProductVariants.ProductTypes"). // Precarregar as variantes do produto
			Preload("Images").
			Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ?", strings.ToLower("%"+searchText+"%"),
				strings.ToLower("%"+searchText+"%")).
			Find(&response.Data).Error; err != nil {
		response.Success = false
		response.Message = "error when searchs products"
		return response, err
	}
	response.Success = true
	response.Message = "Products retrivied successfully"
	return response, nil
}
