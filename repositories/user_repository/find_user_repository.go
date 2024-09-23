package userrepository

import (
	"context"
	"errors"
	"log"

	"github.com/amarantec/e-shop/config/database"
	usermodel "github.com/amarantec/e-shop/models/user_model"
	"gorm.io/gorm"
)

func (r *userRepository) FindUser(ctx context.Context, email string) (usermodel.User, error) {
	userScan := usermodel.User{}

	if err :=
		database.DB.WithContext(ctx).
			Model(&usermodel.User{}).
			Select("id, password").
			Where("email = ?", email).
			Scan(&userScan).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("user repository: user not found -> %v", err)
			return userScan, nil
		}
		log.Printf("user respository error: %v", err)
		return userScan, err
	}

	return userScan, nil
}
