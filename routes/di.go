package routes

import (
	usercontroller "github.com/amarantec/e-shop/controllers/user_controller"
	userrepository "github.com/amarantec/e-shop/repositories/user_repository"
	userservice "github.com/amarantec/e-shop/services/user_service"
)

var userRepo = userrepository.NewUserRepository()
var userSvc = userservice.NewUserService(userRepo)
var userCtrl = usercontroller.NewUserController(userSvc)
