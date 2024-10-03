package cartcontroller

import (
	"context"
	"net/http"
	"time"

	"github.com/amarantec/e-shop/middleware"
	"github.com/gin-gonic/gin"
)

func (ctrl *CartController) GetCartProduct(c *gin.Context) {
	userId := c.GetUint(middleware.UserIdTOKEN)

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := ctrl.service.GetCartProducts(ctxTimeout, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not get cart items count"})
		return
	}

	c.JSON(http.StatusOK, res)
}
