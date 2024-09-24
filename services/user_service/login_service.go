package userservice

import (
	"context"
	"log"

	"github.com/amarantec/e-shop/models"
	usermodel "github.com/amarantec/e-shop/models/user_model"
	"github.com/amarantec/e-shop/utils"
)

func (s *userService) Login(ctx context.Context, user *usermodel.Authenticate) (models.Response[uint], error) {
	response := models.Response[uint]{}

	if validEmail, err := validUserEmail(user.Email); err != nil || !validEmail {
		response.Success = false
		return response, ErrUserEmailInvalidFormat
	}

	if validPassword, err := validUserPassword(user.Password); err != nil || !validPassword {
		response.Success = false
		return response, ErrUserPasswordInvalidFormat
	}

	userRetrivied, err := s.userRepo.FindUserByEmail(ctx, user.Email)
	if err != nil {
		response.Success = false
		return response, err
	}

	passwordIsValid :=
		utils.CheckPassword(user.Password, userRetrivied.Password)
	if !passwordIsValid {
		log.Printf(err.Error())
		response.Success = false
		response.Message = "Wrong password"
		return response, nil
	}

	response.Data = userRetrivied.ID
	response.Success = true
	response.Message = "Successfull login"

	return response, nil
}
