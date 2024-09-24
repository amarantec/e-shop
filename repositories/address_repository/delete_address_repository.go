package addressrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	addressmodel "github.com/amarantec/e-shop/models/address_model"
)

func (r *addressRepository) DeleteAddress(ctx context.Context, userId, addressId uint) (models.Response[bool], error) {
	response := models.Response[bool]{}

	result :=
		database.DB.WithContext(ctx).Where("user_id = ? AND id = ?", userId, addressId).
			Delete(&addressmodel.Address{})

	if result.RowsAffected == 0 {
		response.Success = true
		response.Message = "Address not found"
		return response, nil
	}

	response.Data = true
	response.Success = true
	response.Message = "Address deleted successfuly"
	return response, nil
}
