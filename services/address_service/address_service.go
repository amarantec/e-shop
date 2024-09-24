package addressservice

import (
	"context"
	"errors"
	"unicode"
	"unicode/utf8"

	"github.com/amarantec/e-shop/models"
	addressmodel "github.com/amarantec/e-shop/models/address_model"
	addressrepository "github.com/amarantec/e-shop/repositories/address_repository"
)

type AddressService interface {
	InsertAddress(ctx context.Context, address addressmodel.Address) (models.Response[addressmodel.Address], error)
	DeleteAddress(ctx context.Context, userId, addressId uint) (models.Response[bool], error)
	GetAddress(ctx context.Context, userId, addressId uint) (models.Response[addressmodel.Address], error)
	ListAddresses(ctx context.Context, userId uint) (models.Response[[]addressmodel.Address], error)
	UpdateAddress(ctx context.Context, address addressmodel.Address) (models.Response[bool], error)
}

type addressService struct {
	addressRepo addressrepository.AddressRepository
}

func NewAddressService(repo addressrepository.AddressRepository) AddressService {
	return &addressService{addressRepo: repo}
}

func validAddress(a addressmodel.Address) (bool, error) {
	if a.UserId == 0 {
		return false, ErrAddressUserIdEmpty
	}

	if utf8.RuneCountInString(a.FirstName) < 2 || utf8.RuneCountInString(a.FirstName) > 50 {
		return false, ErrAddressFirstNameInvalidFormat
	}

	if utf8.RuneCountInString(a.LastName) < 2 || utf8.RuneCountInString(a.LastName) > 50 {
		return false, ErrAddressLastNameInvalidFormat
	}

	if utf8.RuneCountInString(a.Street) < 2 || utf8.RuneCountInString(a.Street) > 50 {
		return false, ErrAddressStreetInvalidFormat
	}

	if utf8.RuneCountInString(a.City) < 2 || utf8.RuneCountInString(a.City) > 50 {
		return false, ErrAddressCityInvalidFormat
	}

	if utf8.RuneCountInString(a.State) != 2 {
		return false, ErrAddressStateInvalidFormat
	}

	if utf8.RuneCountInString(a.Zip) != 8 {
		return false, ErrAddressZipInvalidFormat
	}

	for _, char := range a.Zip {
		if !unicode.IsDigit(char) {
			return false, ErrAddressZipInvalidFormat
		}
	}

	if utf8.RuneCountInString(a.Country) != 2 {
		return false, ErrAddressCountryInvalidFormat
	}

	return true, nil
}

var ErrAddressUserIdEmpty = errors.New("address user id is empty")
var ErrAddressStreetInvalidFormat = errors.New("address street invalid format, address street must contain between 3 and 50 chars")
var ErrAddressCityInvalidFormat = errors.New("address city invalid format, address city must contain between 2 and 30 chars")
var ErrAddressStateInvalidFormat = errors.New("address state invalid format, address state must contain only 2 chars. Example -> State: RS")
var ErrAddressCountryInvalidFormat = errors.New("address country invalid format, address country must contain onlye 2 chars. Example -> Country: BR")
var ErrAddressZipInvalidFormat = errors.New("address zip code invalid format, zip must contain 8 digits in range 0-9")
var ErrAddressFirstNameInvalidFormat = errors.New("address first name to short, min 3 characters")
var ErrAddressLastNameInvalidFormat = errors.New("address last name to short, min 3 characters")
var ErrAddressInvalidFormat = errors.New("address invalid format")
var ErrAddressIdEmpty = errors.New("address id is empty")
