package productcontroller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl *ProductController) ListProducts(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := ctrl.service.ListProducts(ctxTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not list this addresses"})
		return
	}

	c.JSON(http.StatusOK, res)
}
