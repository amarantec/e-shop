package productrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (r *productRepository) UpdateProduct(ctx context.Context, product productmodel.Product) (models.Response[bool], error) {
	response := models.Response[bool]{}

	result :=
		database.DB.WithContext(ctx).
			Where("id = ?", product.ID).
			Updates(productmodel.Product{
				Title:       product.Title,
				Description: product.Description,
				ImageURL:    product.ImageURL,
				CategoryID:  product.CategoryID,
				Visible:     product.Visible,
				Featured:    product.Featured,
			})

	if result.Error != nil {
		response.Success = false
		response.Message = "error when updating product register"
		return response, result.Error
	}

	if result.RowsAffected == 0 {
		response.Success = true
		response.Message = "Product not found"
		return response, nil
	}

	response.Data = true
	response.Success = true
	response.Message = "Product updated successfully"
	return response, nil

}
