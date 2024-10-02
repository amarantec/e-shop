package orderservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	ordermodel "github.com/amarantec/e-shop/models/order_model"
	orderrepository "github.com/amarantec/e-shop/repositories/order_repository"
)

type OrderService interface {
	GetOrderDetails(ctx context.Context, userId, orderId uint) (models.Response[ordermodel.OrderDetailsResponse], error)
	GetOrders(ctx context.Context, userId uint) (models.Response[[]ordermodel.OrderOverviewResponse], error)
	PlaceOrder(ctx context.Context, userId uint) (models.Response[bool], error)
}

type orderService struct {
	orderRepo orderrepository.OrderRepository
}

func NewOrderService(repo orderrepository.OrderRepository) OrderService {
	return &orderService{orderRepo: repo}
}
