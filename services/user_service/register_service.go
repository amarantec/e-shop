package userservice

import usermodel "github.com/amarantec/project777/models/user_model"

func (s *userService) Register(user *usermodel.User) error {
	return s.userRepo.Register(user)
}
