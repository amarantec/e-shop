package categoryservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	categorymodel "github.com/amarantec/e-shop/models/category_model"
)

func (s *categoryService) ListCategories(ctx context.Context) (models.Response[[]categorymodel.Category], error) {
	return s.categoryRepo.ListCategories(ctx)
}
