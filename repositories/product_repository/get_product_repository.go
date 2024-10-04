package productrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (r *productRepository) GetProduct(ctx context.Context, productId uint) (productmodel.Product, error) {
	product := productmodel.Product{}
	if err := database.DB.WithContext(ctx).
		Preload("Category").                     // Precarregar a categoria associada
		Preload("ProductVariants.ProductTypes"). // Precarregar as variantes do produto e o ProductTypes
		Preload("Images").                       // Precarregar as imagens do produto
		Where("id = ?", productId).
		First(&product).Error; err != nil {
		return productmodel.Product{}, err
	}

	return product, nil
}
