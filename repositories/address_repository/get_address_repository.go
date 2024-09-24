package addressrepository

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	addressmodel "github.com/amarantec/e-shop/models/address_model"
	"gorm.io/gorm"
)

func (r *addressRepository) GetAddress(ctx context.Context, userId, addressId uint) (models.Response[addressmodel.Address], error) {
	response := models.Response[addressmodel.Address]{}

	if err :=
		database.DB.WithContext(ctx).
			Where("user_id = ? AND id = ?", userId, addressId).
			First(&response.Data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Message = "Address not found"
			return response, nil
		}
		response.Success = false
		return response, err
	}

	response.Success = true
	response.Message = "Address retrivied successfully"
	return response, nil
}
