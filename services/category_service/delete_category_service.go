package categoryservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
)

func (s *categoryService) DeleteCategory(ctx context.Context, categoryId uint) (models.Response[bool], error) {
	return s.categoryRepo.DeleteCategory(ctx, categoryId)
}
