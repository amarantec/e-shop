package categorymodel

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name    string `json:"name" binding:"omitempty" gorm:"not nul"`
	URL     string `json:"url" binding:"omitempty" gorm:"not nul"`
	Visible bool   `json:"visible" binding:"omitempty" gorm:"not nul"`
	Deleted bool   `json:"deleted" binding:"omitempty" gorm:"not nul"`
	Editing bool   `json:"editing" binding:"omitempty" gorm:"not nul"`
	IsNew   bool   `json:"is_new" gorm:"not nul"`
}
