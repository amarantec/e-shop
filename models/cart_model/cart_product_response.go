package cartmodel

type CartProductResponse struct {
	ProductID      uint    `json:"product_id"`
	Title          string  `json:"title"`
	ProductTypesID uint    `json:"product_types_id"`
	ProductTypes   string  `json:"product_types"`
	ImageUrl       string  `json:"image_url"`
	Price          float64 `json:"price"`
	Quantity       int     `json:"quantity"`
}
