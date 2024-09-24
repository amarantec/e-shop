package addressservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
)

func (s *addressService) DeleteAddress(ctx context.Context, userId, addressId uint) (models.Response[bool], error) {
	response := models.Response[bool]{}
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

	response, err := s.addressRepo.DeleteAddress(ctx, userId, addressId)
	if err != nil {
		response.Success = false
		response.Message = "service error: could not delete address"
		return response, err
	}

	return response, nil
}
