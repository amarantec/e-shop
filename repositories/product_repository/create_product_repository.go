package productrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (r *productRepository) CreateProduct(ctx context.Context, product productmodel.Product) (models.Response[productmodel.Product], error) {
	response := models.Response[productmodel.Product]{}

	if err := database.DB.WithContext(ctx).
		Create(&product).Error; err != nil {
		response.Success = false
		response.Message = "error when create product register"
		return response, err
	}

	response.Data = product
	response.Success = true
	response.Message = "Address registred successfully"
	return response, nil
}
