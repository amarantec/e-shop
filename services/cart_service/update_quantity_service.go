package cartservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
)

func (s *cartService) UpdateQuantity(ctx context.Context, cartItem cartmodel.CartItems) (models.Response[bool], error) {
	return s.cartRepo.UpdateQuantity(ctx, cartItem)
}
