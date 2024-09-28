package producttypescontroller

import (
	"context"
	"net/http"
	"time"

	productmodel "github.com/amarantec/e-shop/models/product_model"
	"github.com/gin-gonic/gin"
)

func (ctrl *ProductTypesController) AddProductTypes(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newProductType := productmodel.ProductTypes{}

	if err := c.ShouldBindJSON(&newProductType); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not parse this request"})
		return
	}

	res, err := ctrl.service.AddProductType(ctxTimeout, newProductType)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not add this product type"})
		return
	}

	c.JSON(http.StatusCreated, res)
}
