package productmodel

import "gorm.io/gorm"

type ProductTypes struct {
	gorm.Model
	Name    string `json:"name" binding:"required" gorm:"not nul"`
	Editing bool   `json:"editing"  gorm:"not nul"`
	IsNew   bool   `json:"is_new"  gorm:"not nul"`
}
