package orderrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	ordermodel "github.com/amarantec/e-shop/models/order_model"
)

func (r *orderRepository) GetOrders(ctx context.Context, userId uint) ([]ordermodel.Orders, error) {
	orders := []ordermodel.Orders{}
	if err :=
		database.DB.WithContext(ctx).
			Preload("OrderItems.Product").
			Where("user_id = ?", userId).
			Find(&orders).Error; err != nil {
		return []ordermodel.Orders{}, err
	}

	return orders, nil

}
