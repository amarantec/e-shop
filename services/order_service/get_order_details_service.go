package orderservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	ordermodel "github.com/amarantec/e-shop/models/order_model"
)

func (s orderService) GetOrderDetails(ctx context.Context, userId, orderId uint) (models.Response[ordermodel.OrderDetailsResponse], error) {
	response := models.Response[ordermodel.OrderDetailsResponse]{}

	order, err := s.orderRepo.GetOrderDetails(ctx, userId, orderId)
	if err != nil {
		return response, err
	}

	orderDetailsResponse := ordermodel.OrderDetailsResponse{
		OrderDate:  order.OrderDate,
		TotalPrice: order.TotalPrice,
		Products:   []ordermodel.OrderDetailsProductResponse{},
	}

	for _, item := range order.OrderItems {
		orderDetailsProductResponse := ordermodel.OrderDetailsProductResponse{}
		orderDetailsProductResponse.ProductID = item.ProductID
		orderDetailsProductResponse.ImageURL = item.Product.ImageURL
		orderDetailsProductResponse.ProductType = item.ProductTypes.Name
		orderDetailsProductResponse.Quantity = item.Quantity
		orderDetailsProductResponse.Title = item.Product.Title
		orderDetailsProductResponse.TotalPrice = item.TotalPrice

		orderDetailsResponse.Products = append(orderDetailsResponse.Products, orderDetailsProductResponse)
	}

	response.Data = orderDetailsResponse
	response.Success = true
	return response, nil
}
