package cartservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
)

func (s *cartService) GetCartItemsCount(ctx context.Context, userId uint) (models.Response[int64], error) {
	return s.cartRepo.GetCartItemsCount(ctx, userId)
}
