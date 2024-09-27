package productrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (r *productRepository) ListProductsByCategory(ctx context.Context, categoryUrl string) (models.Response[[]productmodel.Product], error) {
	response := models.Response[[]productmodel.Product]{}
	if err :=
		database.DB.WithContext(ctx).
			Table("products AS p").
			Select(`p.id,
				p.title,
				p.description,
				p.image_url,
				p.category_id,
				COALESCE(p.featured, false) AS featured,
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
			Where("c.url = ?", categoryUrl).
			Scan(&response.Data).Error; err != nil {
		response.Success = false
		return response, err
	}

	response.Success = true
	response.Message = "Products retrivied successfully"
	return response, nil
}
