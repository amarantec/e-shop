package userrepository

import (
	"github.com/amarantec/project777/config/database"
	usermodel "github.com/amarantec/project777/models/user_model"
)

func (r *userRepository) Register(user *usermodel.User) error {
	return database.DB.Create(user).Error
}
