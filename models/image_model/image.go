package imagemodel

import "gorm.io/gorm"

type Images struct {
	gorm.Model
	Data      string `json:"data" binding:"required" gorm:"not nul"`
	ProductID uint
}
