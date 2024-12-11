package request

import "lokajatim/entities"

type TransactionRequest struct {
	UserID      int     `json:"user_id"`
	CartID      int     `json:"cart_id"`
	TotalPrice  float64 `json:"total_price"`
	Status      string  `json:"status"`
	PaymentType string  `json:"payment_type"`
}

func (transactionRequest TransactionRequest) ToEntities() (entities.Transaction, error) {
	transaction := entities.Transaction{
		UserID:      transactionRequest.UserID,
		CartID:      transactionRequest.CartID,
		TotalPrice:  transactionRequest.TotalPrice,
		Status:      transactionRequest.Status,
		PaymentType: transactionRequest.PaymentType,
	}
	return transaction, nil
}
