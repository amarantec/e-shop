package producttyperepository

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/config/database"
	productmodel "github.com/amarantec/e-shop/models/product_model"
	"gorm.io/gorm"
)

func (r *productTypesRepository) GetProductType(ctx context.Context, productTypeId uint) (productmodel.ProductTypes, error) {
	productType := productmodel.ProductTypes{}

	if err :=
		database.DB.WithContext(ctx).
			Where("id = ?", productTypeId).
			First(&productType).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return productmodel.ProductTypes{}, nil
		}
		return productmodel.ProductTypes{}, err
	}

	return productType, nil
}
