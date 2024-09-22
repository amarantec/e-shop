package routes

import (
	usercontroller "github.com/amarantec/project777/controllers/user_controller"
	userrepository "github.com/amarantec/project777/repositories/user_repository"
	userservice "github.com/amarantec/project777/services/user_service"
)

var userRepo = userrepository.NewUserRepository()
var userSvc = userservice.NewUserService(userRepo)
var userCtrl = usercontroller.NewUserController(userSvc)
