package addressrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	addressmodel "github.com/amarantec/e-shop/models/address_model"
)

func (r *addressRepository) InsertAddress(ctx context.Context, address addressmodel.Address) (models.Response[addressmodel.Address], error) {
	response := models.Response[addressmodel.Address]{}
	if err := database.DB.WithContext(ctx).
		Create(&address).Error; err != nil {
		response.Success = false
		response.Message = "error when create address register"
		return response, err
	}

	response.Data = address
	response.Success = true
	response.Message = "Address registred successfully"
	return response, nil
}
