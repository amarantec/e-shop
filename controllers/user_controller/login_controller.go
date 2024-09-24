package usercontroller

import (
	"context"
	"net/http"
	"time"

	usermodel "github.com/amarantec/e-shop/models/user_model"
	"github.com/amarantec/e-shop/utils"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) Login(c *gin.Context) {
	user := usermodel.Authenticate{}
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not parse this request"})
		return
	}

	res, err := ctrl.service.Login(ctxTimeout, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not validate user credentials"})
		return
	}

	token, err := utils.GenerateToken(res.Data, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not generate user token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res, "token": token})
}
