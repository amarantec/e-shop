package ordercontroller

import orderservice "github.com/amarantec/e-shop/services/order_service"

type OrderController struct {
	service orderservice.OrderService
}

func NewOrderController(service orderservice.OrderService) *OrderController {
	return &OrderController{service: service}
}
