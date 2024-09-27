package productrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (r *productRepository) ListProductsByCategory(ctx context.Context, categoryUrl string) (models.Response[[]productmodel.Product], error) {
	response := models.Response[[]productmodel.Product]{}
	if err := database.DB.WithContext(ctx).
		Model(productmodel.Product{}).
		Preload("Category", "url = ?", categoryUrl). // Precarregar a categoria associada
		Preload("ProductVariants.ProductTypes").     // Precarregar as variantes do produto
		Preload("Images").                           // Precarregar as imagens do produto
		Where("visible= ?", true).
		Find(&response.Data).Error; err != nil {
		response.Success = false
		response.Message = "error when getting products"
		return response, err
	}

	response.Success = true
	response.Message = "Products retrieved successfully"
	return response, nil
}
