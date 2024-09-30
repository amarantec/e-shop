package cartrepository

import (
	"context"

	"github.com/amarantec/e-shop/models"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
)

type CartItemsRepository interface {
	GetCartItemsCount(ctx context.Context, userId uint) (models.Response[int64], error)
	ListCartItems(ctx context.Context, userId uint) (models.Response[[]cartmodel.CartProductResponse], error)
	RemoveItemFromCart(ctx context.Context, productID, productTypesID uint) (models.Response[bool], error)
	StoreCartItems(ctx context.Context, cartItems cartmodel.CartItems) (models.Response[cartmodel.CartItems], error)
	UpdateQuantity(ctx context.Context, cartItem cartmodel.CartItems) (models.Response[bool], error)
	GetCartItem(ctx context.Context, cartItem cartmodel.CartItems) (cartmodel.CartItems, error)
}

type cartItemsRepository struct{}

func NewCartItemsRepository() CartItemsRepository {
	return &cartItemsRepository{}
}
