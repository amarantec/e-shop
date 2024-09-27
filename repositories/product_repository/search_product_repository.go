package productrepository

import (
	"context"
	"strings"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (r *productRepository) SearchProducts(ctx context.Context, searchText string) (models.Response[[]productmodel.Product], error) {
	response := models.Response[[]productmodel.Product]{}

	if err :=
		database.DB.WithContext(ctx).
			Table("products AS p").
			Select(`p.id,
				p.title,
				p.description,
				p.image_url,
				p.category_id,
				c.id,
				c.name,
				c.url,
				pv.product_id,
				pv.product_types_id,
				pv.price,
				COALESCE(pv.original_price, 0.0) AS original_price,
				pt.id,
				pt.name`).
			Joins("JOIN categories AS c ON p.category_id = c.id").
			Joins("LEFT JOIN product_variants AS pv ON p.id = pv.product_id").
			Joins("LEFT JOIN product_types AS pt ON pv.product_types_id = pt.id").
			Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ?", strings.ToLower("%"+searchText+"%"),
				strings.ToLower("%"+searchText+"%")).
			Scan(&response.Data).Error; err != nil {
		response.Success = false
		response.Message = "error when get featured products"
		return response, err
	}
	response.Success = true
	response.Message = "Products retrivied successfully"
	return response, nil
}
