package middleware

import (
	"net/http"

	"github.com/amarantec/e-shop/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.ValidateUserToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized,
				gin.H{"error": "authentication required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
