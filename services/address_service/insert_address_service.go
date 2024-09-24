package addressservice

import (
	"context"
	"log"

	"github.com/amarantec/e-shop/models"
	addressmodel "github.com/amarantec/e-shop/models/address_model"
)

func (s *addressService) InsertAddress(ctx context.Context, address addressmodel.Address) (models.Response[addressmodel.Address], error) {
	response := models.Response[addressmodel.Address]{}

	if valid, err := validAddress(address); err != nil || !valid {
		response.Success = false
		return response, ErrAddressInvalidFormat
	}

	response, err := s.addressRepo.InsertAddress(ctx, address)
	if err != nil {
		log.Printf("service error: %v", err)
		response.Success = false
		response.Message = "service error: could not insert this address"
		return response, err
	}

	return response, nil
}
