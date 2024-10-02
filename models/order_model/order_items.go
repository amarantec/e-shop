package ordermodel

import (
	productmodel "github.com/amarantec/e-shop/models/product_model"
	"gorm.io/gorm"
)

type OrderItems struct {
	gorm.Model
	Orders         Orders
	OrdersID       uint `json:"orders_id" binding:"required" gorm:"not null"`
	Product        productmodel.Product
	ProductID      uint `json:"product_id" binding:"required" gorm:"not null"`
	ProductTypes   productmodel.ProductTypes
	ProductTypesID uint    `json:"product_types_id" binding:"required" gorm:"not null"`
	Quantity       int     `json:"quantity" binding:"required" gorm:"not null"`
	TotalPrice     float64 `json:"total_price" gorm:"not null;type:decimal(18,2)"`
}
