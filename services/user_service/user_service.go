package userservice

import (
	"context"
	"errors"
	"regexp"

	"github.com/amarantec/e-shop/models"
	usermodel "github.com/amarantec/e-shop/models/user_model"
	userrepository "github.com/amarantec/e-shop/repositories/user_repository"
)

var ErrUserNameEmpty = errors.New("user name is empty")
var ErrUserNameInvalidFormat = errors.New("user name to short, min 3 characters")
var ErrUserEmailInvalidFormat = errors.New("user email must contain @ and .com")
var ErrUserPasswordInvalidFormat = errors.New("password must contains 8 or more digits")
var ErrUserEmailEmpty = errors.New("user email is empty")
var ErrUserPasswordEmpty = errors.New("user password is empty")

type UserService interface {
	Register(ctx context.Context, user *usermodel.User) (models.Response[uint], error)
	FindUser(ctctx context.Context, user *usermodel.Authenticate) (models.Response[uint], error)
}

type userService struct {
	userRepo userrepository.UserRepository
}

func NewUserService(repo userrepository.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func validUserName(s string) (bool, error) {
	if s == "" {
		return false, ErrUserNameEmpty
	}

	if len(s) < 3 {
		return false, ErrUserNameInvalidFormat
	}

	return true, nil
}

func validUserEmail(s string) (bool, error) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if s == "" {
		return false, ErrUserEmailEmpty
	}

	if !emailRegex.MatchString(s) {
		return false, ErrUserEmailInvalidFormat
	}

	return true, nil
}

func validUserPassword(s string) (bool, error) {
	if s == "" {
		return false, ErrUserPasswordEmpty
	}

	if len(s) < 8 {
		return false, ErrUserPasswordInvalidFormat
	}

	return true, nil
}
