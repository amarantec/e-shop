package wishlistcontroller

import wishlistservice "github.com/amarantec/e-shop/services/wishlist_service"

type WishListController struct {
	service wishlistservice.WishListService
}

func NewWishListController(service wishlistservice.WishListService) *WishListController {
	return &WishListController{service: service}
}
