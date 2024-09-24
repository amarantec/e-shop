package addressrepository

import (
	"context"

	"github.com/amarantec/e-shop/models"
	addressmodel "github.com/amarantec/e-shop/models/address_model"
)

type AddressRepository interface {
	InsertAddress(ctx context.Context, address addressmodel.Address) (models.Response[addressmodel.Address], error)
	ListAddresses(ctx context.Context, userId uint) (models.Response[[]addressmodel.Address], error)
	GetAddress(ctx context.Context, userId, addressId uint) (models.Response[addressmodel.Address], error)
	UpdateAddress(ctx context.Context, address addressmodel.Address) (models.Response[bool], error)
	DeleteAddress(ctx context.Context, userId, addressId uint) (models.Response[bool], error)
}

type addressRepository struct{}

func NewAddressRepository() AddressRepository {
	return &addressRepository{}
}
