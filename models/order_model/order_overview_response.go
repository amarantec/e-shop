package ordermodel

import "time"

type OrderOverviewResponse struct {
	ID              int64     `json:"id"`
	OrderDate       time.Time `json:"order_date"`
	TotalPrice      float64   `json:"total_price"`
	Product         string    `json:"product"`
	ProductImageURL string    `json:"product_image_url"`
}
