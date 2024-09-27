package productservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (s *productService) CreateProduct(ctx context.Context, product productmodel.Product) (models.Response[productmodel.Product], error) {
	response := models.Response[productmodel.Product]{}

	for _, v := range product.ProductVariants {
		v.ProductTypes = productmodel.ProductTypes{}
	}

	if valid, err := validProduct(product); err != nil || !valid {
		response.Success = false
		return response, ErrProductInvalidFormat
	}

	product.Featured = true
	product.Visible = true
	product.Deleted = false
	product.IsNew = true
	product.Editing = false

	response, err := s.productRepo.CreateProduct(ctx, product)
	if err != nil {
		response.Success = false
		return response, err
	}

	return response, nil
}
