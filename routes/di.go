package routes

import (
	addresscontroller "github.com/amarantec/e-shop/controllers/address_controller"
	cartcontroller "github.com/amarantec/e-shop/controllers/cart_controller"
	productcontroller "github.com/amarantec/e-shop/controllers/product_controller"
	producttypescontroller "github.com/amarantec/e-shop/controllers/product_types_controller"
	usercontroller "github.com/amarantec/e-shop/controllers/user_controller"
	addressrepository "github.com/amarantec/e-shop/repositories/address_repository"
	cartrepository "github.com/amarantec/e-shop/repositories/cart_repository"
	productrepository "github.com/amarantec/e-shop/repositories/product_repository"
	producttyperepository "github.com/amarantec/e-shop/repositories/product_type_repository"
	userrepository "github.com/amarantec/e-shop/repositories/user_repository"
	addressservice "github.com/amarantec/e-shop/services/address_service"
	cartservice "github.com/amarantec/e-shop/services/cart_service"
	productservice "github.com/amarantec/e-shop/services/product_service"
	producttypeservice "github.com/amarantec/e-shop/services/product_type_service"
	userservice "github.com/amarantec/e-shop/services/user_service"
)

var userRepo = userrepository.NewUserRepository()
var userSvc = userservice.NewUserService(userRepo)
var userCtrl = usercontroller.NewUserController(userSvc)

var addressRepo = addressrepository.NewAddressRepository()
var addressSvc = addressservice.NewAddressService(addressRepo)
var addressCtrl = addresscontroller.NewAddressController(addressSvc)

var productRepo = productrepository.NewProductRepository()
var productSvc = productservice.NewProductService(productRepo)
var productCtrl = productcontroller.NewProductController(productSvc)

var productTypeRepo = producttyperepository.NewProductTypeRepository()
var productTypeSvc = producttypeservice.NewProductTypesService(productTypeRepo)
var productTypeCtrl = producttypescontroller.NewProductTypesController(productTypeSvc)

var cartRepo = cartrepository.NewCartItemsRepository()
var cartSvc = cartservice.NewCartService(cartRepo)
var cartCtrl = cartcontroller.NewCartController(cartSvc)
