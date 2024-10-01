package categoryrepository

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	categorymodel "github.com/amarantec/e-shop/models/category_model"
	"gorm.io/gorm"
)

func (r *categoryRepository) UpdateCategory(ctx context.Context, category categorymodel.Category) (models.Response[bool], error) {
	response := models.Response[bool]{}

	if err :=
		database.DB.WithContext(ctx).Model(&categorymodel.Category{}).
			Where("id = ?", category.ID).
			Updates(categorymodel.Category{
				Name: category.Name,
				URL:  category.URL,
			}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Message = "Category not found"
			return response, nil
		}
		response.Success = false
		return response, err
	}

	response.Data = true
	response.Success = true
	response.Message = "Category updated successfully"
	return response, nil
}
