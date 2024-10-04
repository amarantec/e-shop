package wishlistmodel

import (
	"gorm.io/gorm"
)

type WishList struct {
	gorm.Model
	UserID         uint `json:"user_id" binding:"required" gorm:"not null"`
	ProductID      uint `json:"product_id" binding:"required" gorm:"not null"`
	ProductTypesID uint `json:"product_types_id" binding:"required" gorm:"not null"`
}
