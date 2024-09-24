package addresscontroller

import addressservice "github.com/amarantec/e-shop/services/address_service"

type AddressController struct {
	service addressservice.AddressService
}

func NewAddressController(service addressservice.AddressService) *AddressController {
	return &AddressController{service: service}
}
