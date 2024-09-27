package productmodel

import "gorm.io/gorm"

type ProductTypes struct {
	gorm.Model
	Name    string `json:"name" binding:"required" gorm:"not nul"`
	Editing bool   `json:"editing" binding:"required" gorm:"not nul"`
	IsNew   bool   `json:"is_new" binding:"required" gorm:"not nul"`
}
