package categorycontroller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl *CategoryController) ListCategories(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := ctrl.service.ListCategories(ctxTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not get list of categories"})
		return
	}

	c.JSON(http.StatusOK, res)
}
