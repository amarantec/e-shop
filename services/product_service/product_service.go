package productservice

import (
	"context"
	"errors"
	"unicode/utf8"

	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
	productrepository "github.com/amarantec/e-shop/repositories/product_repository"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product productmodel.Product) (models.Response[productmodel.Product], error)
	DeleteProduct(ctx context.Context, productId uint) (models.Response[bool], error)
	ListFeaturedProducts(ctx context.Context) (models.Response[[]productmodel.Product], error)
	GetProduct(ctx context.Context, productId uint) (models.Response[productmodel.Product], error)
	ListProducts(ctx context.Context) (models.Response[[]productmodel.Product], error)
	ListProductsByCategory(ctx context.Context, categoryUrl string) (models.Response[[]productmodel.Product], error)
	UpdateProduct(ctx context.Context, product productmodel.Product) (models.Response[bool], error)
	SearchProducts(ctx context.Context, searchText string) (models.Response[[]productmodel.Product], error)
}

type productService struct {
	productRepo productrepository.ProductRepository
}

func NewProductService(repo productrepository.ProductRepository) ProductService {
	return &productService{productRepo: repo}
}

func validProduct(p productmodel.Product) (bool, error) {
	if utf8.RuneCountInString(p.Title) < 2 {
		return false, ErrProductTitleInvalidFormat
	}

	if utf8.RuneCountInString(p.Description) < 5 || utf8.RuneCountInString(p.Description) > 100 {
		return false, ErrProductDescriptionInvalidFormat
	}

	if p.ImageURL == "" {
		return false, ErrProductImageUrlEmpty
	}

	if p.CategoryID <= 0 {
		return false, ErrProductCategoryIdEmpty
	}

	return true, nil
}

var ErrProductTitleInvalidFormat = errors.New("product title must contain 2 or more characters")
var ErrProductDescriptionInvalidFormat = errors.New("product description must contain between 2 and 100 characters")
var ErrProductImageUrlEmpty = errors.New("product image url is empty")
var ErrProductCategoryIdEmpty = errors.New("product category id is empty")
var ErrProductInvalidFormat = errors.New("product invalid format, check logs for more information")
var ErrProductIdEmpty = errors.New("product id is empty")
var ErrProductCategoryUrlEmpty = errors.New("product category url is empty")
var ErrProductSearchTextEmpty = errors.New("product search text is empty")