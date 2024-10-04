package producttyperepository

import (
	"context"

	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

type ProductTypesRepository interface {
	AddProductType(ctx context.Context, productType productmodel.ProductTypes) (models.Response[productmodel.ProductTypes], error)
	ListProductTypes(ctx context.Context) (models.Response[[]productmodel.ProductTypes], error)
	UpdateProductType(ctx context.Context, productType productmodel.ProductTypes) (models.Response[bool], error)
	GetProductType(ctx context.Context, productTypeId uint) (productmodel.ProductTypes, error)
}

type productTypesRepository struct{}

func NewProductTypeRepository() ProductTypesRepository {
	return &productTypesRepository{}
}
