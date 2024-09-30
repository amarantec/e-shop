package cartcontroller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/amarantec/e-shop/middleware"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
	"github.com/gin-gonic/gin"
)

func (ctrl *CartController) StoreCartItems(c *gin.Context) {
	userId := c.GetUint(middleware.UserIdTOKEN)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newCartItem := cartmodel.CartItems{
		UserID: userId,
	}

	if err := c.ShouldBindJSON(&newCartItem); err != nil {
		log.Printf("ctrl error: %v", err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not parse this request"})
		return
	}

	fmt.Println("cart item: ", newCartItem)

	res, err := ctrl.service.StoreCartItems(ctxTimeout, newCartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not store these cart items"})
		return
	}

	c.JSON(http.StatusCreated, res)
}
