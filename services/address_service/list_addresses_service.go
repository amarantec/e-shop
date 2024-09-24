package addressservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	addressmodel "github.com/amarantec/e-shop/models/address_model"
)

func (s *addressService) ListAddresses(ctx context.Context, userId uint) (models.Response[[]addressmodel.Address], error) {
	response := models.Response[[]addressmodel.Address]{}

	if userId <= 0 {
		response.Success = false
		response.Message = "Address user id is empty"
		return response, ErrAddressUserIdEmpty
	}

	response, err := s.addressRepo.ListAddresses(ctx, userId)
	if err != nil {
		response.Success = false
		response.Message = "service error: could not list addresses"
		return response, err
	}

	return response, nil
}
