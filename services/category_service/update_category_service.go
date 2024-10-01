package categoryservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	categorymodel "github.com/amarantec/e-shop/models/category_model"
)

func (s *categoryService) UpdateCategory(ctx context.Context, category categorymodel.Category) (models.Response[bool], error) {
	return s.categoryRepo.UpdateCategory(ctx, category)
}
