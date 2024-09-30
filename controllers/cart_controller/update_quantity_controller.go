package cartcontroller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/amarantec/e-shop/middleware"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
	"github.com/gin-gonic/gin"
)

func (ctrl *CartController) UpdateQuantity(c *gin.Context) {
	userId := c.GetUint(middleware.UserIdTOKEN)

	productIdStr := c.Param("productId")

	productTypeIdStr := c.Param("productTypeId")

	quantityStr := c.Param("quantity")

	productId, err1 := strconv.ParseUint(productIdStr, 0, 64)
	productTypeId, err2 := strconv.ParseUint(productTypeIdStr, 0, 64)
	quantity, err3 := strconv.ParseInt(quantityStr, 10, 32)

	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "invalid parameters"})
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateCartItem := cartmodel.CartItems{
		UserID:         userId,
		ProductID:      uint(productId),
		ProductTypesID: uint(productTypeId),
		Quantity:       int(quantity),
	}

	res, err := ctrl.service.UpdateQuantity(ctxTimeout, updateCartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not update this cart item"})
		return
	}

	c.JSON(http.StatusNoContent, res)
}
