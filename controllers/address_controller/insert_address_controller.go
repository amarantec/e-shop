package addresscontroller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/amarantec/e-shop/middleware"
	addressmodel "github.com/amarantec/e-shop/models/address_model"
	"github.com/gin-gonic/gin"
)

func (ctrl *AddressController) InsertAddress(c *gin.Context) {
	userId := c.GetUint(middleware.UserIdTOKEN)
	var newAddress addressmodel.Address

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Print("user id: %d", userId)

	if err := c.ShouldBindJSON(&newAddress); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not parse this request"})
		return
	}

	newAddress.UserId = userId
	_, err := ctrl.service.InsertAddress(ctxTimeout, newAddress)
	if err != nil {
		fmt.Printf("address controller error: %v", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not insert this address"})
		return
	}

	c.JSON(http.StatusCreated, newAddress.UserId)
}
