package ordercontroller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/amarantec/e-shop/middleware"
	"github.com/gin-gonic/gin"
)

func (ctrl *OrderController) PlaceOrder(c *gin.Context) {
	userId := c.GetUint(middleware.UserIdTOKEN)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := ctrl.service.PlaceOrder(ctxTimeout, userId)
	if err != nil {
		log.Printf("order controller error: %v", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not get orders list"})
		return
	}

	c.JSON(http.StatusOK, res)

}
