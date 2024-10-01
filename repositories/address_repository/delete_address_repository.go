package addressrepository

import (
	"context"
	"errors"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	addressmodel "github.com/amarantec/e-shop/models/address_model"
	"gorm.io/gorm"
)

func (r *addressRepository) DeleteAddress(ctx context.Context, userId, addressId uint) (models.Response[bool], error) {
	response := models.Response[bool]{}

	if err :=
		database.DB.WithContext(ctx).
			Where("user_id = ? AND id = ?", userId, addressId).
			Unscoped().Delete(&addressmodel.Address{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Message = "Product not found"
			return response, nil
		}
		response.Success = false
		return response, err
	}

	response.Data = true
	response.Success = true
	response.Message = "Address deleted successfuly"
	return response, nil
}
