package userrepository

import (
	"context"
	"fmt"

	"github.com/amarantec/e-shop/config/database"
	usermodel "github.com/amarantec/e-shop/models/user_model"
)

func (r *userRepository) ChangePassword(ctx context.Context, userId uint, password string) error {
	result :=
		database.DB.WithContext(ctx).Model(&usermodel.User{}).
			Where("id = ?", userId).
			Updates(usermodel.User{
				Password: password,
			})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
