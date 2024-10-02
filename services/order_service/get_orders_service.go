package orderservice

import (
	"context"
	"fmt"

	"github.com/amarantec/e-shop/models"
	ordermodel "github.com/amarantec/e-shop/models/order_model"
)

func (s orderService) GetOrders(ctx context.Context, userId uint) (models.Response[[]ordermodel.OrderOverviewResponse], error) {
	response := models.Response[[]ordermodel.OrderOverviewResponse]{}

	orders, err := s.orderRepo.GetOrders(ctx, userId)
	if err != nil {
		return response, err
	}

	orderResponse := []ordermodel.OrderOverviewResponse{}
	for _, item := range orders {
		o := ordermodel.OrderOverviewResponse{}
		o.ID = int64(item.ID)
		o.OrderDate = item.OrderDate
		o.TotalPrice = item.TotalPrice

		if len(item.OrderItems) > 1 {
			o.Product = fmt.Sprintf("%s and %d more...", item.OrderItems[0].Product.Title, len(item.OrderItems)-1)
		} else if len(item.OrderItems) == 1 {
			o.Product = item.OrderItems[0].Product.Title
		} else {
			o.Product = "No products"
		}

		if len(item.OrderItems) > 0 {
			o.ProductImageURL = item.OrderItems[0].Product.ImageURL
		}

		orderResponse = append(orderResponse, o)
	}

	response.Data = orderResponse
	response.Success = true
	return response, nil
}
