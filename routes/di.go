package routes

import (
	addresscontroller "github.com/amarantec/e-shop/controllers/address_controller"
	usercontroller "github.com/amarantec/e-shop/controllers/user_controller"
	addressrepository "github.com/amarantec/e-shop/repositories/address_repository"
	userrepository "github.com/amarantec/e-shop/repositories/user_repository"
	addressservice "github.com/amarantec/e-shop/services/address_service"
	userservice "github.com/amarantec/e-shop/services/user_service"
)

var userRepo = userrepository.NewUserRepository()
var userSvc = userservice.NewUserService(userRepo)
var userCtrl = usercontroller.NewUserController(userSvc)

var addressRepo = addressrepository.NewAddressRepository()
var addressSvc = addressservice.NewAddressService(addressRepo)
var addressCtrl = addresscontroller.NewAddressController(addressSvc)
