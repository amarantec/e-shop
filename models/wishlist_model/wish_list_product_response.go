package wishlistmodel

type WishListProductResponse struct {
	ProductID   uint   `json:"product_id"`
	Title       string `json:"title"`
	ProductType string `json:"product_type"`
	ImageURL    string `json:"image_url"`
}
