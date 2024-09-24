package addressservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	addressmodel "github.com/amarantec/e-shop/models/address_model"
)

func (s *addressService) UpdateAddress(ctx context.Context, address addressmodel.Address) (models.Response[bool], error) {
	response := models.Response[bool]{}

	if valid, err := validAddress(address); err != nil || !valid {
		response.Success = false
		return response, ErrAddressInvalidFormat
	}

	response, err := s.addressRepo.UpdateAddress(ctx, address)
	if err != nil {
		response.Success = false
		response.Message = "service error: could not update this address"
		return response, err
	}

	return response, nil
}
