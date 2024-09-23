package usermodel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"  binding:"required" gorm:"not null"`
	Email    string `json:"email" binding:"required" gorm:"not null; unique"`
	Password string `json:"password" binding:"required" gorm:"not null"`
}
