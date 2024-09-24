package addresscontroller

import (
	"context"
	"net/http"
	"time"

	"github.com/amarantec/e-shop/middleware"
	addressmodel "github.com/amarantec/e-shop/models/address_model"
	"github.com/gin-gonic/gin"
)

func (ctrl *AddressController) UpdateAddress(c *gin.Context) {
	userId := c.GetUint(middleware.UserIdTOKEN)
	newAddress := addressmodel.Address{
		UserId: userId,
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&newAddress); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not parse this request"})
		return
	}

	res, err := ctrl.service.UpdateAddress(ctxTimeout, newAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not update this address"})
		return
	}

	c.JSON(http.StatusOK, res)
}
