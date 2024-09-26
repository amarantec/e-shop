package productmodel

import (
	categorymodel "github.com/amarantec/e-shop/models/category_model"
	imagemodel "github.com/amarantec/e-shop/models/image_model"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title           string              `json:"title" binding:"required" gorm:"not nul"`
	Description     string              `json:"description" binding:"required" gorm:"not nul"`
	ImageURL        string              `json:"image_url" binding:"required" gorm:"not nul"`
	Images          []imagemodel.Images `gorm:"foreignKey:ProductID"`
	Category        categorymodel.Category
	CategoryID      uint              `json:"category_id" binding:"required" gorm:"not nul"`
	Featured        bool              `json:"featured" binding:"omitempty" gorm:"not nul"`
	ProductVariants []ProductVariants `gorm:"foreignKey:ProductID"`
	Visible         bool              `json:"visible" binding:"omitempty" gorm:"not nul"`
	Deleted         bool              `json:"deleted" binding:"omitempty" gorm:"not nul"`
	Editing         bool              `json:"editing" binding:"omitempty" gorm:"not nul"`
	IsNew           bool              `json:"is_new" gorm:"not nul, omitempty"`
}
