package cartservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
)

func (s *cartService) RemoveItemFromCart(ctx context.Context, productID, productTypesID uint) (models.Response[bool], error) {
	return s.cartRepo.RemoveItemFromCart(ctx, productID, productTypesID)
}
