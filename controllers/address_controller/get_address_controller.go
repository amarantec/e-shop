package addresscontroller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/amarantec/e-shop/middleware"
	"github.com/gin-gonic/gin"
)

func (ctrl *AddressController) GetAddress(c *gin.Context) {
	userId := c.GetUint(middleware.UserIdTOKEN)

	addressIdStr := c.Param("addressId")
	addressId, err := strconv.Atoi(addressIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "invalid paramter"})
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := ctrl.service.GetAddress(ctxTimeout, userId, uint(addressId))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not get this address"})
		return
	}

	c.JSON(http.StatusOK, res)
}
