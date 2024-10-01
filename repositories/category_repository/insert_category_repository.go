package categoryrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	categorymodel "github.com/amarantec/e-shop/models/category_model"
)

func (r *categoryRepository) InsertCategory(ctx context.Context, category categorymodel.Category) (models.Response[categorymodel.Category], error) {
	response := models.Response[categorymodel.Category]{}

	if err :=
		database.DB.WithContext(ctx).
			Create(&category).Error; err != nil {
		response.Message = "could not create this category"
		return response, err
	}

	response.Data = category
	response.Success = true
	response.Message = "Category successfully created"
	return response, nil
}
