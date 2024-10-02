package orderrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	ordermodel "github.com/amarantec/e-shop/models/order_model"
)

func (r *orderRepository) PlaceOrder(ctx context.Context, order ordermodel.Orders) (models.Response[bool], error) {
	response := models.Response[bool]{}

	if err :=
		database.DB.WithContext(ctx).
			Create(&order).Error; err != nil {
		return response, err
	}

	response.Data = true
	response.Success = true
	response.Message = "Order placed successfully"
	return response, nil
}
