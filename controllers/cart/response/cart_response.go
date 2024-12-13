package response

import (
	"lokajatim/entities"
	"time"
)

// CartResponse is the response for the cart controller
// @Description CartResponse is the response for cart data retrieval
// @Param ID int true "ID of the cart"
// @Param UserID int true "ID of the user"
// @Param User UserResponse true "User of the cart"
// @Param Items []CartItemResponse true "Items in the cart"
// @Param TotalPrice float64 true "Total price of the cart"
// @Param TotalPriceAfterTransaction float64 true "Total price of the cart after transaction"
// @Param CreatedAt string true "Created at of the cart"
// @Param UpdatedAt string true "Updated at of the cart"
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
// @Description CartItemResponse is the response for cart item data retrieval
// @Param ID int true "ID of the cart item"
// @Param ProductID int true "ID of the product"
// @Param Product ProductResponse true "Product of the cart item"
// @Param Quantity int true "Quantity of the product"
// @Param CreatedAt string true "Created at of the cart item"
// @Param UpdatedAt string true "Updated at of the cart item"
type CartItemResponse struct {
	ID        int             `json:"id"`
	ProductID int             `json:"product_id"`
	Product   ProductResponse `json:"product"`
	Quantity  int             `json:"quantity"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

// UserResponse is the response for the user controller
// @Description UserResponse is the response for user data retrieval
// @Param ID int true "ID of the user"
// @Param Name string true "Name of the user"
// @Param Email string true "Email of the user"
// @Param Address string true "Address of the user"
// @Param PhoneNumber string true "Phone number of the user"
type UserResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

// ProductResponse is the response for the product controller
// @Description ProductResponse is the response for product data retrieval
// @Param ID int true "ID of the product"
// @Param Name string true "Name of the product"
// @Param Price int true "Price of the product"
// @Param Stock int true "Stock of the product"
// @Param Description string true "Description of the product"
// @Param CreatedAt string true "Created at of the product"
// @Param UpdatedAt string true "Updated at of the product"
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
