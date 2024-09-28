package producttypescontroller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl *ProductTypesController) ListProductTypes(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := ctrl.service.ListProductTypes(ctxTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not list product types"})
		return
	}

	c.JSON(http.StatusOK, res)
}
