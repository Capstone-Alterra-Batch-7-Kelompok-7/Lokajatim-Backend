package response

import (
	"lokajatim/entities"
	"time"
)

// CartResponse is the response for the cart controller
type CartResponse struct {
	ID                         int                `json:"id"`
	UserID                     int                `json:"user_id"`
	User                       UserResponse       `json:"user"`
	Items                      []CartItemResponse `json:"items"`
	TotalPrice                 float64            `json:"total_price"`
	TotalPriceAfterTransaction float64            `json:"total_price_after_transaction"`
	CreatedAt                  time.Time          `json:"created_at"`
	UpdatedAt                  time.Time          `json:"updated_at"`
}

// CartItemResponse is the response for the cart item controller
type CartItemResponse struct {
	ID        int             `json:"id"`
	ProductID int             `json:"product_id"`
	Product   ProductResponse `json:"product"`
	Photos    []string        `json:"photos"`
	Quantity  int             `json:"quantity"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

// UserResponse is the response for the user controller
type UserResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

// ProductResponse is the response for the product controller
type ProductResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	Photos      []string  `json:"photos"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CartFromEntities converts a Cart entity to CartResponse
func CartFromEntities(cart entities.Cart) CartResponse {
	userResponse := UserResponse{
		ID:          cart.User.ID,
		Name:        cart.User.Name,
		Email:       cart.User.Email,
		Address:     cart.User.Address,
		PhoneNumber: cart.User.PhoneNumber,
	}

	var itemResponses []CartItemResponse
	for _, item := range cart.Items {
		photoUrls := make([]string, len(item.Product.Photos))
		for i, photo := range item.Product.Photos {
			photoUrls[i] = photo.UrlPhoto
		}

		itemResponses = append(itemResponses, CartItemResponse{
			ID:        item.ID,
			ProductID: item.ProductID,
			Product:   ProductFromEntities(item.Product),
			Photos:    photoUrls,
			Quantity:  item.Quantity,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return CartResponse{
		ID:                         cart.ID,
		UserID:                     cart.UserID,
		User:                       userResponse,
		Items:                      itemResponses,
		TotalPrice:                 cart.TotalPrice,
		TotalPriceAfterTransaction: cart.TotalPriceAfterTransaction,
		CreatedAt:                  cart.CreatedAt,
		UpdatedAt:                  cart.UpdatedAt,
	}
}

// ProductFromEntities converts a Product entity to ProductResponse
func ProductFromEntities(product entities.Product) ProductResponse {
	photoUrls := make([]string, len(product.Photos))
	for i, photo := range product.Photos {
		photoUrls[i] = photo.UrlPhoto
	}

	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.Description,
		Photos:      photoUrls,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
