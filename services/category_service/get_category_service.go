package categoryservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	categorymodel "github.com/amarantec/e-shop/models/category_model"
)

func (s *categoryService) GetCategory(ctx context.Context, categoryId uint) (models.Response[categorymodel.Category], error) {
	return s.categoryRepo.GetCategory(ctx, categoryId)
}
