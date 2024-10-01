package categorycontroller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	categorymodel "github.com/amarantec/e-shop/models/category_model"
	"github.com/gin-gonic/gin"
)

func (ctrl *CategoryController) UpdateCategory(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	categoryIdStr := c.Param("categoryId")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "invalid parameter"})
		return
	}

	updateCategory := categorymodel.Category{}

	if err := c.ShouldBindJSON(&updateCategory); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not decode this request"})
		return
	}

	updateCategory.ID = uint(categoryId)

	res, err := ctrl.service.UpdateCategory(ctxTimeout, updateCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not get category"})
		return
	}

	c.JSON(http.StatusNoContent, res)
}
