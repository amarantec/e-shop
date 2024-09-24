package addressrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	addressmodel "github.com/amarantec/e-shop/models/address_model"
)

func (r *addressRepository) UpdateAddress(ctx context.Context, address addressmodel.Address) (models.Response[bool], error) {
	response := models.Response[bool]{}
	result :=
		database.DB.WithContext(ctx).Model(&addressmodel.Address{}).
			Where("user_id = ? AND id = ?", address.UserId, address.ID).
			Updates(addressmodel.Address{
				FirstName: address.FirstName,
				LastName:  address.LastName,
				Street:    address.Street,
				City:      address.City,
				State:     address.State,
				Zip:       address.Zip,
				Country:   address.Country,
			})

	if result.Error != nil {
		response.Success = false
		response.Message = "error when updating address register"
		return response, result.Error
	}

	if result.RowsAffected == 0 {
		response.Success = true
		response.Message = "Address not found"
		return response, nil
	}

	response.Data = true
	response.Success = true
	response.Message = "Address updated successfully"
	return response, nil

}
