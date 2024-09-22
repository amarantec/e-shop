package usercontroller

import userservice "github.com/amarantec/project777/services/user_service"

type UserController struct {
	service userservice.UserService
}

func NewUserController(service userservice.UserService) *UserController {
	return &UserController{service: service}
}
