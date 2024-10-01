package categorycontroller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl *CategoryController) DeleteCategory(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	categoryIdStr := c.Param("categoryId")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "invalid parameter"})
		return
	}

	res, err := ctrl.service.DeleteCategory(ctxTimeout, uint(categoryId))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not delete category"})
		return
	}

	c.JSON(http.StatusNoContent, res)
}
