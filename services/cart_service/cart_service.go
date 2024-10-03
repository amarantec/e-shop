package cartservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
	cartrepository "github.com/amarantec/e-shop/repositories/cart_repository"
)

type CartService interface {
	GetCartProducts(ctx context.Context, userId uint) (models.Response[[]cartmodel.CartProductResponse], error)
	StoreCartItems(ctx context.Context, cartItems cartmodel.CartItems) (models.Response[cartmodel.CartItems], error)
	GetCartItemsCount(ctx context.Context, userId uint) (models.Response[int64], error)
	UpdateQuantity(ctx context.Context, cartItem cartmodel.CartItems) (models.Response[bool], error)
	RemoveItemFromCart(ctx context.Context, productID, productTypesID uint) (models.Response[bool], error)
}

type cartService struct {
	cartRepo cartrepository.CartItemsRepository
}

func NewCartService(repo cartrepository.CartItemsRepository) CartService {
	return &cartService{cartRepo: repo}
}
