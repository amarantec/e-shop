package main

import (
	"log"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	database.ConnectDatabase()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
