package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", userCtrl.Register)
	}
}
