package cartcontroller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl *CartController) RemoveItemFromCart(c *gin.Context) {
	productIdStr := c.Param("productId")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "invalid paramter"})
		return
	}

	productTypesIdStr := c.Param("productTypesId")
	productTypesId, err := strconv.Atoi(productTypesIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "invalid paramter"})
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := ctrl.service.RemoveItemFromCart(ctxTimeout, uint(productId), uint(productTypesId))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not remove this item"})
		return
	}

	c.JSON(http.StatusNoContent, res)
}
