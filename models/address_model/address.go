package addressmodel

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserId    uint   `json:"user_id" binding:"required" gorm:"not null"`
	FirstName string `json:"first_name" binding:"required" gorm:"not null"`
	LastName  string `json:"last_name" binding:"required" gorm:"not null"`
	Street    string `json:"street" binding:"required" gorm:"not null"`
	City      string `json:"city" binding:"required" gorm:"not null"`
	State     string `json:"state" binding:"required" gorm:"not null"`
	Zip       string `json:"zip" binding:"required" gorm:"not null"`
	Country   string `json:"country" binding:"required" gorm:"not null"`
}
