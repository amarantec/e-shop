package categoryrepository

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	categorymodel "github.com/amarantec/e-shop/models/category_model"
	"gorm.io/gorm"
)

func (r *categoryRepository) DeleteCategory(ctx context.Context, categoryId uint) (models.Response[bool], error) {
	response := models.Response[bool]{}

	if err :=
		database.DB.WithContext(ctx).Model(&categorymodel.Category{}).
			Where("id= ?", categoryId).
			Unscoped().Delete(&categorymodel.Category{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Message = "Category not found"
			return response, nil
		}
		response.Success = false
		return response, err
	}

	response.Data = true
	response.Success = true
	response.Message = "Category deleted successfully"
	return response, nil
}
