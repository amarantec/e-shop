package productcontroller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	productmodel "github.com/amarantec/e-shop/models/product_model"
	"github.com/gin-gonic/gin"
)

func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateProduct := productmodel.Product{}

	productIdStr := c.Param("productId")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "invalid parameter"})
		return
	}

	updateProduct.ID = uint(productId)

	if err := c.ShouldBindJSON(&updateProduct); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not parse this request"})
		return
	}

	res, err := ctrl.service.UpdateProduct(ctxTimeout, updateProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not update this product"})
		return
	}
	c.JSON(http.StatusNoContent, res)
}
