package main

import (
	"github.com/amarantec/project777/config/database"
	"github.com/amarantec/project777/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	routes.SetupRoutes()

	r.Run(":8080")
}