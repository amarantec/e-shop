package categorycontroller

import categoryservice "github.com/amarantec/e-shop/services/category_service"

type CategoryController struct {
	service categoryservice.CategoryService
}

func NewCategoryController(service categoryservice.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}
