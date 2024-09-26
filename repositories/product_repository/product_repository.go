package productrepository

import (
	"context"

	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product productmodel.Product) (models.Response[productmodel.Product], error)
	DeleteProduct(ctx context.Context, productId uint) (models.Response[bool], error)
	ListFeaturedProducts(ctx context.Context) (models.Response[[]productmodel.Product], error)
	GetProduct(ctx context.Context, productId uint) (models.Response[productmodel.Product], error)
	ListProducts(ctx context.Context) (models.Response[[]productmodel.Product], error)
	ListProductsByCategory(ctx context.Context, categoryUrl string) (models.Response[[]productmodel.Product], error)
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}
