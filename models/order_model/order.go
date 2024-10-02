package ordermodel

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	UserID     uint      `json:"user_id" gorm:"not null"`
	OrderDate  time.Time `json:"order_date" gorm:"not null"`
	TotalPrice float64   `json:"total_price" gorm:"not null"`
	OrderItems []OrderItems
}
