package productrepository

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
	"gorm.io/gorm"
)

func (r *productRepository) DeleteProduct(ctx context.Context, productId uint) (models.Response[bool], error) {
	response := models.Response[bool]{}

	if err := database.DB.WithContext(ctx).Where("id = ?", productId).
		Delete(&productmodel.Product{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Message = "Product not found"
			return response, nil
		}
		response.Success = false
		return response, err
	}

	response.Data = true
	response.Success = true
	response.Message = "Product deleted successfuly"
	return response, nil
}
