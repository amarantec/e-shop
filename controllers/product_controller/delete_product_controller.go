package productcontroller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl *ProductController) DeleteProduct(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	productIdStr := c.Param("productId")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "invalid paramter"})
		return
	}

	res, err := ctrl.service.DeleteProduct(ctxTimeout, uint(productId))

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not delete this product"})
		return
	}

	c.JSON(http.StatusNoContent, res)
}
