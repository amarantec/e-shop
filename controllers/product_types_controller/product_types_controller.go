package producttypescontroller

import producttypeservice "github.com/amarantec/e-shop/services/product_type_service"

type ProductTypesController struct {
	service producttypeservice.ProductTypesService
}

func NewProductTypesController(service producttypeservice.ProductTypesService) *ProductTypesController {
	return &ProductTypesController{service: service}
}
