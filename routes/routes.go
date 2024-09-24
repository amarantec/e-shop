package routes

import (
	"github.com/amarantec/e-shop/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", userCtrl.SaveUser)
		userGroup.POST("/login", userCtrl.Login)
		userGroup.POST("/change-password/:password", middleware.Auth, userCtrl.ChangePassword)
	}
}
