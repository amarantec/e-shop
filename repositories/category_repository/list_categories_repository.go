package categoryrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	categorymodel "github.com/amarantec/e-shop/models/category_model"
)

func (r *categoryRepository) ListCategories(ctx context.Context) (models.Response[[]categorymodel.Category], error) {
	response := models.Response[[]categorymodel.Category]{}

	if err :=
		database.DB.WithContext(ctx).
			Find(&response.Data).Error; err != nil {
		response.Message = "could not get list of categories"
		return response, err
	}

	response.Success = true
	response.Message = "Category list successfully retrivied"
	return response, nil
}
