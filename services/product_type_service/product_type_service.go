package producttypeservice

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
	producttyperepository "github.com/amarantec/e-shop/repositories/product_type_repository"
)

type ProductTypesService interface {
	AddProductType(ctx context.Context, productTypes productmodel.ProductTypes) (models.Response[productmodel.ProductTypes], error)
	ListProductTypes(ctx context.Context) (models.Response[[]productmodel.ProductTypes], error)
	UpdateProductType(ctx context.Context, productTypes productmodel.ProductTypes) (models.Response[bool], error)
}

type productTypesService struct {
	productTypesRepo producttyperepository.ProductTypesRepository
}

func NewProductTypesService(repo producttyperepository.ProductTypesRepository) ProductTypesService {
	return &productTypesService{productTypesRepo: repo}
}

var ErrProductTypesIdEmpty = errors.New("product types id is empty")
var ErrProductTypesNameEmpty = errors.New("product types name is empty")
