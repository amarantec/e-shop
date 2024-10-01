package categorycontroller

import (
	"context"
	"net/http"
	"time"

	categorymodel "github.com/amarantec/e-shop/models/category_model"
	"github.com/gin-gonic/gin"
)

func (ctlr *CategoryController) InsertCategory(c *gin.Context) {
	category := categorymodel.Category{}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not decode this request"})
		return
	}

	res, err := ctlr.service.InsertCategory(ctxTimeout, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not insert category"})
		return
	}

	c.JSON(http.StatusCreated, res)
}
