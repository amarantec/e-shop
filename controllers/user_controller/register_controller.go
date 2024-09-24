package usercontroller

import (
	"context"
	"net/http"
	"time"

	usermodel "github.com/amarantec/e-shop/models/user_model"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) SaveUser(c *gin.Context) {
	var newUser usermodel.User
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse this request"})
		return
	}

	res, err := ctrl.service.SaveUser(ctxTimeout, &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not register user"})
		return
	}

	c.JSON(http.StatusCreated, res)
}
