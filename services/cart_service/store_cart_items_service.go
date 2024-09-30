package cartservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
)

func (s *cartService) StoreCartItems(ctx context.Context, cartItems cartmodel.CartItems) (models.Response[cartmodel.CartItems], error) {
	response := models.Response[cartmodel.CartItems]{}

	cart, err := s.cartRepo.GetCartItem(ctx, cartItems)
	if err != nil {
		return response, err
	}

	if cart.UserID == cartItems.UserID && cart.ProductID == cartItems.ProductID && cart.ProductTypesID == cartItems.ProductTypesID {
		cart.Quantity = cartItems.Quantity
		s.cartRepo.UpdateQuantity(ctx, cart)
		response.Data = cart
		response.Success = true
		response.Message = "Quantity updated"
		return response, nil
	}

	return s.cartRepo.StoreCartItems(ctx, cartItems)

}
