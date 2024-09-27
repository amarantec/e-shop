package productcontroller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl *ProductController) GetProduct(c *gin.Context) {
	productIdStr := c.Param("productId")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "invalid parameter"})
		return
	}
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := ctrl.service.GetProduct(ctxTimeout, uint(productId))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not get this product"})
		return
	}

	c.JSON(http.StatusOK, res)
}
