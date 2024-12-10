package response

import (
	"lokajatim/entities"
	"time"
)

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

type CartItemResponse struct {
	ID        int             `json:"id"`
	ProductID int             `json:"product_id"`
	Product   ProductResponse `json:"product"`
	Quantity  int             `json:"quantity"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type UserResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type ProductResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

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
		itemResponses = append(itemResponses, CartItemResponse{
			ID:        item.ID,
			ProductID: item.ProductID,
			Product:   ProductFromEntities(item.Product),
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

func ProductFromEntities(product entities.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
