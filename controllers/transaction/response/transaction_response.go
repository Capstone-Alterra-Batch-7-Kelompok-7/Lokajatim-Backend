package response

import (
	"lokajatim/entities"
	"time"
)

type TransactionResponse struct {
	ID            int          `json:"id"`
	TransactionID string       `json:"transaction_id"`
	User          UserResponse `json:"user"`
	CartID        int          `json:"cart_id"`
	TotalPrice    float64      `json:"total_price"`
	Status        string       `json:"status"`
	PaymentURL    string       `json:"payment_url"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
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

func TransactionFromEntity(transaction entities.Transaction) TransactionResponse {

	userResponse := UserResponse{
		ID:          transaction.User.ID,
		Name:        transaction.User.Name,
		Email:       transaction.User.Email,
		Address:     transaction.User.Address,
		PhoneNumber: transaction.User.PhoneNumber,
	}

	return TransactionResponse{
		ID:            transaction.ID,
		TransactionID: transaction.TransactionID,
		User:          userResponse,
		CartID:        transaction.CartID,
		TotalPrice:    transaction.TotalPrice,
		Status:        transaction.Status,
		PaymentURL:    transaction.PaymentURL,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}
