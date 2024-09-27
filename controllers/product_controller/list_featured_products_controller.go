package productcontroller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl *ProductController) ListFeaturedProducts(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := ctrl.service.ListFeaturedProducts(ctxTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not get featured products list"})
		return
	}

	c.JSON(http.StatusOK, res)
}
