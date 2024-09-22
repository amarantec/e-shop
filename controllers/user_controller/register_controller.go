package usercontroller

import (
	"net/http"

	usermodel "github.com/amarantec/project777/models/user_model"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) Register(c *gin.Context) {
	var newUser usermodel.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.Register(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Could not register user"})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}
