package routes

import (
	addresscontroller "github.com/amarantec/e-shop/controllers/address_controller"
	cartcontroller "github.com/amarantec/e-shop/controllers/cart_controller"
	categorycontroller "github.com/amarantec/e-shop/controllers/category_controller"
	ordercontroller "github.com/amarantec/e-shop/controllers/order_controller"
	productcontroller "github.com/amarantec/e-shop/controllers/product_controller"
	producttypescontroller "github.com/amarantec/e-shop/controllers/product_types_controller"
	usercontroller "github.com/amarantec/e-shop/controllers/user_controller"
	wishlistcontroller "github.com/amarantec/e-shop/controllers/wishlist_controller"
	addressrepository "github.com/amarantec/e-shop/repositories/address_repository"
	cartrepository "github.com/amarantec/e-shop/repositories/cart_repository"
	categoryrepository "github.com/amarantec/e-shop/repositories/category_repository"
	orderrepository "github.com/amarantec/e-shop/repositories/order_repository"
	productrepository "github.com/amarantec/e-shop/repositories/product_repository"
	producttyperepository "github.com/amarantec/e-shop/repositories/product_type_repository"
	userrepository "github.com/amarantec/e-shop/repositories/user_repository"
	wishlistrepository "github.com/amarantec/e-shop/repositories/wishlist_repository"
	addressservice "github.com/amarantec/e-shop/services/address_service"
	cartservice "github.com/amarantec/e-shop/services/cart_service"
	categoryservice "github.com/amarantec/e-shop/services/category_service"
	orderservice "github.com/amarantec/e-shop/services/order_service"
	productservice "github.com/amarantec/e-shop/services/product_service"
	producttypeservice "github.com/amarantec/e-shop/services/product_type_service"
	userservice "github.com/amarantec/e-shop/services/user_service"
	wishlistservice "github.com/amarantec/e-shop/services/wishlist_service"
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

var categoryRepo = categoryrepository.NewCategoryRepository()
var categorySvc = categoryservice.NewCategoryService(categoryRepo)
var categoryCtrl = categorycontroller.NewCategoryController(categorySvc)

var orderRepo = orderrepository.NewOrderRepository()
var orderSvc = orderservice.NewOrderService(orderRepo)
var orderCtrl = ordercontroller.NewOrderController(orderSvc)

var wishListRepo = wishlistrepository.NewWishListRepository()
var wishListSvc = wishlistservice.NewWishListService(wishListRepo)
var wishListCtrl = wishlistcontroller.NewWishListController(wishListSvc)
