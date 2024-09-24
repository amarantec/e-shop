package addressrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	addressmodel "github.com/amarantec/e-shop/models/address_model"
)

func (r *addressRepository) ListAddresses(ctx context.Context, userId uint) (models.Response[[]addressmodel.Address], error) {
	response := models.Response[[]addressmodel.Address]{}

	if err :=
		database.DB.WithContext(ctx).
			Where("user_id = ?", userId).Find(&response.Data).Error; err != nil {
		response.Success = false
		response.Message = "error when searching addresses in database"
		return response, err
	}

	response.Success = true
	response.Message = "Addresses retrivied successfuly"
	return response, nil
}
