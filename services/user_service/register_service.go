package userservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	usermodel "github.com/amarantec/e-shop/models/user_model"
	"github.com/amarantec/e-shop/utils"
)

func (s *userService) SaveUser(ctx context.Context, user *usermodel.User) (models.Response[uint], error) {
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

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		response.Success = false
		response.Message = "error when hashing password"
		return response, err
	}

	user.Password = hashedPassword

	response, err = s.userRepo.SaveUser(ctx, user)
	if err != nil {
		response.Success = false
		return response, err
	}

	return response, nil
}
