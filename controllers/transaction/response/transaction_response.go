package response

import (
	"lokajatim/entities"
	"time"
)

type TransactionResponse struct {
	ID            int           `json:"id"`
	TransactionID string        `json:"transaction_id"`
	UserID        int           `json:"user_id"`
	User          entities.User `json:"user"`
	CartID        int           `json:"cart_id"`
	Cart          entities.Cart `json:"cart"`
	TotalPrice    float64       `json:"total_price"`
	Status        string        `json:"status"`
	PaymentType   string        `json:"payment_type"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

func TransactionFromEntity(transaction entities.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:            transaction.ID,
		TransactionID: transaction.TransactionID,
		UserID:        transaction.UserID,
		User:          transaction.User,
		CartID:        transaction.CartID,
		Cart:          transaction.Cart,
		TotalPrice:    transaction.TotalPrice,
		Status:        transaction.Status,
		PaymentType:   transaction.PaymentType,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}
