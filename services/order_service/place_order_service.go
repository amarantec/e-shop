package orderservice

import (
	"context"
	"log"
	"time"

	"github.com/amarantec/e-shop/models"
	ordermodel "github.com/amarantec/e-shop/models/order_model"
	cartrepository "github.com/amarantec/e-shop/repositories/cart_repository"
)

func (s orderService) PlaceOrder(ctx context.Context, userId uint) (models.Response[bool], error) {
	response := models.Response[bool]{}
	products, err := cartrepository.NewCartItemsRepository().ListCartItems(ctx, userId)
	if err != nil {
		log.Printf("place order error: %v", err)
		response.Message = "could not get list of items"
		return response, err
	}

	totalPrice := 0.0
	for _, item := range products.Data {
		totalPrice += (item.Price * float64(item.Quantity))
	}

	orderItems := []ordermodel.OrderItems{}

	for _, item := range products.Data {
		oi := ordermodel.OrderItems{}

		oi.ProductID = item.ProductID
		oi.ProductTypesID = item.ProductTypesID
		oi.Quantity = item.Quantity
		oi.TotalPrice = item.Price * float64(item.Quantity)

		orderItems = append(orderItems, oi)
	}

	order := ordermodel.Orders{
		UserID:     userId,
		OrderDate:  time.Now(),
		TotalPrice: totalPrice,
		OrderItems: orderItems,
	}

	response, err = s.orderRepo.PlaceOrder(ctx, order)
	if err != nil {
		return response, err
	}

	if deleted, err := cartrepository.NewCartItemsRepository().DeleteCart(ctx, userId); err != nil || !deleted {
		return response, err
	}

	response.Data = true
	response.Success = true
	return response, nil
}
