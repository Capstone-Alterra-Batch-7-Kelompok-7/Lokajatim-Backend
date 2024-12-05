package response

import (
	"lokajatim/entities"
	"time"
)

// CartResponse is the response for the Cart controller

type CartResponse struct {
	ID        int                `json:"id"`
	UserID    int                `json:"user_id"`
	User      UserResponse       `json:"user"`
	Items     []CartItemResponse `json:"items"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
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
	ID   int    `json:"id"`
	Name string `json:"name"`
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
		ID:   cart.User.ID,
		Name: cart.User.Name,
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
		ID:        cart.ID,
		UserID:    cart.UserID,
		User:      userResponse,
		Items:     itemResponses,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
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
