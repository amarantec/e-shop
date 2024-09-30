package cartmodel

import "gorm.io/gorm"

type CartItems struct {
	gorm.Model
	UserID         uint `json:"user_id" binding:"required" gorm:"not null"`
	ProductID      uint `json:"product_id" binding:"required" gorm:"not null"`
	ProductTypesID uint `json:"product_types_id" binding:"required" gorm:"not null"`
	Quantity       int  `json:"quantity" binding:"required" gorm:"not null"`
}
