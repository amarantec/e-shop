package userservice

import (
	"context"

	"github.com/amarantec/e-shop/models"
	"github.com/amarantec/e-shop/utils"
)

func (s *userService) ChangeUserPassword(ctx context.Context, userId uint, password string) (models.Response[bool], error) {
	response := models.Response[bool]{}

	if userId <= 0 {
		response.Success = false
		response.Message = "Invalid user id"
		return response, ErrUserIdEmpty
	}

	if validPassword, err := validUserPassword(password); err != nil || !validPassword {
		response.Success = false
		response.Message = "Invalid password format"
		return response, ErrUserPasswordInvalidFormat
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		response.Success = false
		response.Message = "error when hashing password"
		return response, err
	}

	if err := s.userRepo.ChangePassword(ctx, userId, hashedPassword); err != nil {
		response.Success = false
		response.Message = "error updating user password"
		return response, err
	}

	response.Data = true
	response.Success = true
	response.Message = "Password updated successfully"
	return response, nil
}
