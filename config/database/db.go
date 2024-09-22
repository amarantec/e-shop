package database

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test,db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
}
