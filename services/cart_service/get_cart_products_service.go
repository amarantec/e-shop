package cartservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
)

func (s *cartService) GetCartProduct(ctx context.Context, userId uint) (models.Response[[]cartmodel.CartProductResponse], error) {
	return s.cartRepo.ListCartItems(ctx, userId)
}
