package middleware

import (
	"net/http"

	"github.com/amarantec/e-shop/utils"
	"github.com/gin-gonic/gin"
)

const UserIdTOKEN = "UserId"

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is empty"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	c.Set(UserIdTOKEN, userId)
	c.Next()
}
