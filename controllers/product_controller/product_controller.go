package productcontroller

import productservice "github.com/amarantec/e-shop/services/product_service"

type ProductController struct {
	service productservice.ProductService
}

func NewProductController(service productservice.ProductService) *ProductController {
	return &ProductController{service: service}
}
