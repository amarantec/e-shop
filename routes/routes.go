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

	authenticated := r.Group("/address")

	authenticated.POST("/insert-address", middleware.Auth, addressCtrl.InsertAddress)
	authenticated.DELETE("/delete-address/:addressId", middleware.Auth, addressCtrl.DeleteAddress)
	authenticated.GET("/list-addresses", middleware.Auth, addressCtrl.ListAddresses)
	authenticated.GET("/get-address/:addressId", middleware.Auth, addressCtrl.GetAddress)
	authenticated.PUT("/update-address", middleware.Auth, addressCtrl.UpdateAddress)

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
}
