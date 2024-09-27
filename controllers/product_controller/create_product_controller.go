package productcontroller

import (
	"context"
	"log"
	"net/http"
	"time"

	productmodel "github.com/amarantec/e-shop/models/product_model"
	"github.com/gin-gonic/gin"
)

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	newProduct := productmodel.Product{}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		log.Printf("p controller error: %v", err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not parse this request"})
		return
	}

	res, err := ctrl.service.CreateProduct(ctxTimeout, newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not register this product"})
		return
	}

	c.JSON(http.StatusCreated, res)
}
