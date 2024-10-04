package categoryservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	categorymodel "github.com/amarantec/e-shop/models/category_model"
)

func (s *categoryService) InsertCategory(ctx context.Context, category categorymodel.Category) (models.Response[categorymodel.Category], error) {
	category.Visible = true
	category.IsNew = true

	return s.categoryRepo.InsertCategory(ctx, category)
}
