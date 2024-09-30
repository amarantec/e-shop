package cartrepository

import (
	"context"

	"github.com/amarantec/e-shop/config/database"
	"github.com/amarantec/e-shop/models"
	cartmodel "github.com/amarantec/e-shop/models/cart_model"
	productmodel "github.com/amarantec/e-shop/models/product_model"
)

func (r *cartItemsRepository) ListCartItems(ctx context.Context, userId uint) (models.Response[[]cartmodel.CartProductResponse], error) {
	response := models.Response[[]cartmodel.CartProductResponse]{}
	cartItems := []cartmodel.CartItems{}
	products := productmodel.Product{}
	variants := productmodel.ProductVariants{}
	if err :=
		database.DB.WithContext(ctx).
			Where("user_id = ?", userId).
			Find(&cartItems).Error; err != nil {
		response.Success = false
		return response, err
	}

	for _, item := range cartItems {
		if err :=
			database.DB.WithContext(ctx).
				Preload("ProductVariants.ProductTypes").
				Where("id = ?", item.ProductID).
				Find(&products).Error; err != nil {
			response.Success = false
			return response, err
		}

		if err :=
			database.DB.WithContext(ctx).
				Preload("ProductTypes").
				Where("product_id = ? AND product_types_id = ?", item.ProductID, item.ProductTypesID).
				Find(&variants).Error; err != nil {
			response.Success = false
			return response, err
		}

		cartProduct := cartmodel.CartProductResponse{
			ProductID:      products.ID,
			Title:          products.Title,
			ProductTypesID: item.ProductTypesID,
			ProductTypes:   variants.ProductTypes.Name,
			ImageUrl:       products.ImageURL,
			Price:          variants.Price,
			Quantity:       item.Quantity,
		}

		response.Data = append(response.Data, cartProduct)
	}

	response.Success = true
	response.Message = "Cart Product successfully retrivied"
	return response, nil
}
