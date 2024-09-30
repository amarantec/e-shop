package cartcontroller

import cartservice "github.com/amarantec/e-shop/services/cart_service"

type CartController struct {
	service cartservice.CartService
}

func NewCartController(service cartservice.CartService) *CartController {
	return &CartController{service: service}
}
