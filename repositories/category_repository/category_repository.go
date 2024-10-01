package categoryrepository

import (
	"context"

	"github.com/amarantec/e-shop/models"
	categorymodel "github.com/amarantec/e-shop/models/category_model"
)

type CategoryRepository interface {
	ListCategories(ctx context.Context) (models.Response[[]categorymodel.Category], error)
	UpdateCategory(ctx context.Context, category categorymodel.Category) (models.Response[bool], error)
	InsertCategory(ctx context.Context, category categorymodel.Category) (models.Response[categorymodel.Category], error)
	GetCategory(ctx context.Context, categoryId uint) (models.Response[categorymodel.Category], error)
	DeleteCategory(ctx context.Context, categoryId uint) (models.Response[bool], error)
}

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}
