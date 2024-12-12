package request

import "lokajatim/entities"

type TransactionRequest struct {
	UserID     int     `json:"user_id"`
	CartID     int     `json:"cart_id"`
}

func (transactionRequest TransactionRequest) ToEntities() (entities.Transaction, error) {
	transaction := entities.Transaction{
		UserID:     transactionRequest.UserID,
		CartID:     transactionRequest.CartID,
	}
	return transaction, nil
}
