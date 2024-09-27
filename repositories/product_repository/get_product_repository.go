package productrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (r *productRepository) GetProduct(ctx context.Context, productId uint) (models.Response[productmodel.Product], error) {
	response := models.Response[productmodel.Product]{}

	if err := database.DB.WithContext(ctx).
		Model(productmodel.Product{}).
		Preload("Category").                     // Precarregar a categoria associada
		Preload("ProductVariants.ProductTypes"). // Precarregar as variantes do produto e o ProductTypes
		Preload("Images").                       // Precarregar as imagens do produto
		Where("id= ?", productId).
		First(&response.Data).Error; err != nil {
		response.Success = false
		response.Message = "error when getting product"
		return response, err
	}

	response.Success = true
	response.Message = "Product retrieved successfully"
	return response, nil
}
