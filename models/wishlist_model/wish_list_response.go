package wishlistmodel

import (
	"gorm.io/gorm"
)

type WishListResponse struct {
	gorm.Model
	QuantityItems int64
	Products      []WishListProductResponse
}
