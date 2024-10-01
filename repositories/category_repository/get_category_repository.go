package categoryrepository

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	categorymodel "github.com/amarantec/e-shop/models/category_model"
	"gorm.io/gorm"
)

func (r *categoryRepository) GetCategory(ctx context.Context, categoryId uint) (models.Response[categorymodel.Category], error) {
	response := models.Response[categorymodel.Category]{}

	if err :=
		database.DB.WithContext(ctx).
			Model(&categorymodel.Category{}).
			Where("id = ?", categoryId).
			First(&response.Data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Message = "Category not found"
			return response, nil
		}
		response.Success = false
		return response, err
	}

	response.Success = true
	response.Message = "Category successfully retrivied"
	return response, nil
}
