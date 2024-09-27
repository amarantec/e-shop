package database

import (
	"fmt"
	"os"

	addressmodel "github.com/amarantec/e-shop/models/address_model"
	categorymodel "github.com/amarantec/e-shop/models/category_model"
	imagemodel "github.com/amarantec/e-shop/models/image_model"
	productmodel "github.com/amarantec/e-shop/models/product_model"
	usermodel "github.com/amarantec/e-shop/models/user_model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s  user=%s  password=%s dbname=%s port=%s sslmode=disable  TimeZone=America/Sao_Paulo", host, username, password, databaseName, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(
		&usermodel.User{},
		&addressmodel.Address{},
		&categorymodel.Category{},
		&productmodel.Product{},
		&productmodel.ProductTypes{},
		&productmodel.ProductVariants{},
		&imagemodel.Images{},
	)
}
