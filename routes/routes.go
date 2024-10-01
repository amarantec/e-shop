package routes

import (
	"github.com/amarantec/e-shop/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", userCtrl.SaveUser)
		userGroup.POST("/login", userCtrl.Login)
		userGroup.POST("/change-password/:password", middleware.Auth, userCtrl.ChangePassword)
	}

	addressGroup := r.Group("/address")
	{
		addressGroup.POST("/insert-address", middleware.Auth, addressCtrl.InsertAddress)
		addressGroup.DELETE("/delete-address/:addressId", middleware.Auth, addressCtrl.DeleteAddress)
		addressGroup.GET("/list-addresses", middleware.Auth, addressCtrl.ListAddresses)
		addressGroup.GET("/get-address/:addressId", middleware.Auth, addressCtrl.GetAddress)
		addressGroup.PUT("/update-address", middleware.Auth, addressCtrl.UpdateAddress)
	}

	categoryGroup := r.Group("/category")
	{
		categoryGroup.POST("insert-category", categoryCtrl.InsertCategory)
		categoryGroup.GET("/list-categories", categoryCtrl.ListCategories)
		categoryGroup.GET("/get-category/:categoryId", categoryCtrl.GetCategory)
		categoryGroup.PUT("/update-category/:categoryId", categoryCtrl.UpdateCategory)
		categoryGroup.DELETE("/delete-category/:categoryId", categoryCtrl.DeleteCategory)
	}

	productGroup := r.Group("/products")
	{
		productGroup.POST("/create-product", productCtrl.CreateProduct)
		productGroup.GET("/get-product/:productId", productCtrl.GetProduct)
		productGroup.GET("/list-products", productCtrl.ListProducts)
		productGroup.GET("/list-featured-products", productCtrl.ListFeaturedProducts)
		productGroup.GET("/list-products-by-category/:categoryUrl", productCtrl.ListProductsByCategory)
		productGroup.GET("/search-products/:searchText", productCtrl.SearchProducts)
		productGroup.DELETE("/delete-product/:productId", productCtrl.DeleteProduct)
		productGroup.PUT("/update-product/:productId", productCtrl.UpdateProduct)
	}

	productTypesGroup := r.Group("/product-types")
	{
		productTypesGroup.POST("/add-product-type", productTypeCtrl.AddProductTypes)
		productTypesGroup.GET("/list-product-types", productTypeCtrl.ListProductTypes)
		productTypesGroup.PUT("/update-product-type/:productTypeId", productTypeCtrl.UpdateProductType)
	}

	cartItemsGroup := r.Group("/cart")
	{
		cartItemsGroup.POST("/store-cart-items", middleware.Auth, cartCtrl.StoreCartItems)
		cartItemsGroup.GET("/get-cart-product", middleware.Auth, cartCtrl.GetCartProduct)
		cartItemsGroup.GET("/get-cart-items-count", middleware.Auth, cartCtrl.GetCartItemsCount)
		cartItemsGroup.DELETE("/remove-item-from-cart/:productId/:productTypesId", middleware.Auth, cartCtrl.RemoveItemFromCart)
		cartItemsGroup.PUT("/update-quantity/:productId/:productTypeId/:quantity", middleware.Auth, cartCtrl.UpdateQuantity)
	}
}
