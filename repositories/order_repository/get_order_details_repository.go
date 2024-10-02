package orderrepository

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/config/database"
	ordermodel "github.com/amarantec/e-shop/models/order_model"
	"gorm.io/gorm"
)

func (r *orderRepository) GetOrderDetails(ctx context.Context, userId, orderId uint) (ordermodel.Orders, error) {
	order := ordermodel.Orders{}
	if err :=
		database.DB.WithContext(ctx).
			Preload("OrderItems.Product").
			Preload("OrderItems.ProductTypes").
			Where("user_id = ? AND id = ?", userId, orderId).
			First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ordermodel.Orders{}, nil
		}
		return ordermodel.Orders{}, err
	}

	return order, nil
}
