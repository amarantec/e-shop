package usermodel

type Authenticate struct {
	Email    string `json:"email" binding:"required" gorm:"not null; unique"`
	Password string `json:"password" binding:"required" gorm:"not null"`
}
