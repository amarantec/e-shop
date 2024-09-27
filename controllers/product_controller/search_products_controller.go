package productcontroller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl *ProductController) SearchProducts(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	searchText := c.Param("searchText")

	res, err := ctrl.service.SearchProducts(ctxTimeout, searchText)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not search products"})
		return
	}
	c.JSON(http.StatusOK, res)
}
