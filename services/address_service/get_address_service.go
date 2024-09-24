package addressservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	addressmodel "github.com/amarantec/e-shop/models/address_model"
)

func (s *addressService) GetAddress(ctx context.Context, userId, addressId uint) (models.Response[addressmodel.Address], error) {
	response := models.Response[addressmodel.Address]{}
	if userId <= 0 {
		response.Success = false
		response.Message = "User id is empty"
		return response, ErrAddressUserIdEmpty
	}

	if addressId <= 0 {
		response.Success = false
		response.Message = "Address id is empty"
		return response, ErrAddressIdEmpty
	}

	response, err := s.addressRepo.GetAddress(ctx, userId, addressId)
	if err != nil {
		response.Success = false
		response.Message = "service error: could not search address"
		return response, err
	}

	return response, nil
}
