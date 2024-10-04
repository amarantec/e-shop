package wishlistcontroller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/amarantec/e-shop/middleware"
	"github.com/gin-gonic/gin"
)

func (ctrl *WishListController) RemoveFromWishList(c *gin.Context) {
	userId := c.GetUint(middleware.UserIdTOKEN)
	itemId, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "invalid parameter"})
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := ctrl.service.RemoveFromWishList(ctxTimeout, userId, uint(itemId))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not remove this item from wish list"})
		return
	}

	c.JSON(http.StatusOK, res)
}
