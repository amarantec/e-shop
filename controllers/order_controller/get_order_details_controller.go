package ordercontroller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/amarantec/e-shop/middleware"
	"github.com/gin-gonic/gin"
)

func (ctrl *OrderController) GetOrderDetails(c *gin.Context) {
	userId := c.GetUint(middleware.UserIdTOKEN)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	orderIdStr := c.Param("orderId")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "invalid parameters"})
		return
	}

	res, err := ctrl.service.GetOrderDetails(ctxTimeout, userId, uint(orderId))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not get order details list"})
		return
	}

	c.JSON(http.StatusOK, res)
}
