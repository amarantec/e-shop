package productmodel

import (
	"gorm.io/gorm"
)

type ProductVariants struct {
	gorm.Model
	ProductID      uint         `json:"product_id" binding:"required" gorm:"not nul"`
	ProductTypes   ProductTypes 
	ProductTypesID uint         `json:"product_types_id" binding:"required" gorm:"not nul"`
	Price          float64      `json:"price" binding:"required" gorm:"not nul"`
	OriginalPrice  float64      `json:"original_price" binding:"required" gorm:"not nul"`
	Visible        bool         `json:"visible" binding:"omitempty" gorm:"not nul"`
	Deleted        bool         `json:"deleted" binding:"omitempty" gorm:"not nul"`
	Editing        bool         `json:"editing" binding:"omitempty" gorm:"not nul"`
	IsNew          bool         `json:"is_new"  gorm:"not nul"`
}
