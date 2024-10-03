package orderservice

import (
	"context"
	"time"

	"github.com/amarantec/e-shop/models"
	ordermodel "github.com/amarantec/e-shop/models/order_model"
	cartrepository "github.com/amarantec/e-shop/repositories/cart_repository"
	productrepository "github.com/amarantec/e-shop/repositories/product_repository"
)

func (s orderService) PlaceOrder(ctx context.Context, userId uint) (models.Response[bool], error) {
	response := models.Response[bool]{}
	cartItems, err := cartrepository.NewCartItemsRepository().GetCartProducts(ctx, userId)
	if err != nil {
		return response, err
	}

	orderItems := []ordermodel.OrderItems{}
	totalPrice := 0.0
	for _, item := range cartItems {
		product, _ := productrepository.NewProductRepository().GetProduct(ctx, item.ProductID)
		for _, v := range product.ProductVariants {
			totalPrice += (v.Price * float64(item.Quantity))
			oi := ordermodel.OrderItems{}
			oi.ProductID = item.ProductID
			oi.ProductTypesID = item.ProductTypesID
			oi.Quantity = item.Quantity
			oi.TotalPrice = v.Price * float64(item.Quantity)
			orderItems = append(orderItems, oi)
		}
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
