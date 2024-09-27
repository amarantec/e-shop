package productcontroller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl *ProductController) ListProductsByCategory(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	categoryUrl := c.Param("categoryUrl")

	res, err := ctrl.service.ListProductsByCategory(ctxTimeout, categoryUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not list products"})
		return
	}
	c.JSON(http.StatusOK, res)
}
