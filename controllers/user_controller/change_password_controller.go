package usercontroller

import (
	"context"
	"net/http"
	"time"

	"github.com/amarantec/e-shop/middleware"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) ChangePassword(c *gin.Context) {
	userId := c.GetUint(middleware.UserIdTOKEN)

	password := c.Param("password")
	if password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter"})
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := ctrl.service.ChangeUserPassword(ctxTimeout, userId, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not change password"})
		return
	}

	c.JSON(http.StatusOK, res)
}
