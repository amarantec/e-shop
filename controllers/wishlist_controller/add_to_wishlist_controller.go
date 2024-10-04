package wishlistcontroller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/amarantec/e-shop/middleware"
	wishlistmodel "github.com/amarantec/e-shop/models/wishlist_model"
	"github.com/gin-gonic/gin"
)

func (ctrl *WishListController) AddToWishList(c *gin.Context) {
	userId := c.GetUint(middleware.UserIdTOKEN)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	item := wishlistmodel.WishList{
		UserID: userId,
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		log.Printf("add to wish list error: %v", err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not parse this request"})
		return
	}

	res, err := ctrl.service.AddToWishList(ctxTimeout, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not add this item to wish list"})
		return
	}

	c.JSON(http.StatusCreated, res)
}
