package userservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	usermodel "github.com/amarantec/e-shop/models/user_model"
)

func (s *userService) Register(ctx context.Context, user *usermodel.User) (models.Response[uint], error) {
	response := models.Response[uint]{}

	if validName, err := validUserName(user.Name); err != nil || !validName {
		response.Success = false
		return response, ErrUserNameInvalidFormat
	}

	if validEmail, err := validUserEmail(user.Email); err != nil || !validEmail {
		response.Success = false
		return response, ErrUserEmailInvalidFormat
	}

	if validPassword, err := validUserPassword(user.Password); err != nil || !validPassword {
		response.Success = false
		return response, ErrUserPasswordInvalidFormat
	}

	response, err := s.userRepo.Register(ctx, user)
	if err != nil {
		response.Success = false
		return response, err
	}

	return response, nil
}
