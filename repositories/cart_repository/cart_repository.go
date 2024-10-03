package cartrepository

import (
	"context"

	"github.com/amarantec/e-shop/models"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
)

type CartItemsRepository interface {
	GetCartProducts(ctx context.Context, userId uint) ([]cartmodel.CartItems, error)
	GetCartItemsCount(ctx context.Context, userId uint) (models.Response[int64], error)
	RemoveItemFromCart(ctx context.Context, productID, productTypesID uint) (models.Response[bool], error)
	StoreCartItems(ctx context.Context, cartItems cartmodel.CartItems) (models.Response[cartmodel.CartItems], error)
	UpdateQuantity(ctx context.Context, cartItem cartmodel.CartItems) (models.Response[bool], error)
	GetCartItem(ctx context.Context, cartItem cartmodel.CartItems) (cartmodel.CartItems, error)
	DeleteCart(ctx context.Context, userId uint) (bool, error)
}

type cartItemsRepository struct{}

func NewCartItemsRepository() CartItemsRepository {
	return &cartItemsRepository{}
}
