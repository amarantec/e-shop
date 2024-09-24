package usermodel

import (
	addressmodel "github.com/amarantec/e-shop/models/address_model"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string                 `json:"name"  binding:"required" gorm:"not null"`
	Email    string                 `json:"email" binding:"required" gorm:"not null; unique"`
	Password string                 `json:"password" binding:"required" gorm:"not null"`
	Address  []addressmodel.Address `json:"address" gorm:"foreignKey:UserId"`
}
