package userrepository

import (
	"context"
	"errors"
	"log"

	"github.com/amarantec/e-shop/config/database"
	usermodel "github.com/amarantec/e-shop/models/user_model"
	"gorm.io/gorm"
)

func (r *userRepository) FindUserById(ctx context.Context, userId uint) (usermodel.User, error) {
	userScan := usermodel.User{}

	if err :=
		database.DB.WithContext(ctx).
			Where("id = ?", userId).
			First(&userScan).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("user repository: user not found -> %v", err)
			return userScan, nil
		}
		log.Printf("user repository error: %v", err)
		return userScan, err
	}

	return userScan, nil
}
