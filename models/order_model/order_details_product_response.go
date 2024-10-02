package ordermodel

type OrderDetailsProductResponse struct {
	ProductID   uint    `json:"product_id"`
	Title       string  `json:"title"`
	ProductType string  `json:"product_type"`
	ImageURL    string  `json:"image_url"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}
