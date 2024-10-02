package orderrepository

import (
	"context"

	"github.com/amarantec/e-shop/models"
	ordermodel "github.com/amarantec/e-shop/models/order_model"
)

type OrderRepository interface {
	PlaceOrder(ctx context.Context, order ordermodel.Orders) (models.Response[bool], error)
	GetOrders(ctx context.Context, userId uint) ([]ordermodel.Orders, error)
	GetOrderDetails(ctx context.Context, userId, orderId uint) (ordermodel.Orders, error)
}

type orderRepository struct{}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}
