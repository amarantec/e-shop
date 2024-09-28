package producttypescontroller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	productmodel "github.com/amarantec/e-shop/models/product_model"
	"github.com/gin-gonic/gin"
)

func (ctrl *ProductTypesController) UpdateProductType(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateProductType := productmodel.ProductTypes{}

	productTypeIdStr := c.Param("productTypeId")
	productTypeId, err := strconv.Atoi(productTypeIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "invalid parameter"})
		return
	}

	updateProductType.ID = uint(productTypeId)

	if err := c.ShouldBindJSON(&updateProductType); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not parse this request"})
		return
	}

	res, err := ctrl.service.UpdateProductType(ctxTimeout, updateProductType)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not update this product type"})
		return
	}

	c.JSON(http.StatusOK, res)
}
