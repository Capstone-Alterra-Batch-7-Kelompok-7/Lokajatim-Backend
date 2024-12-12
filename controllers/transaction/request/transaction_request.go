package request

import "lokajatim/entities"

// TransactionRequest is the request for the transaction controller
// @Description TransactionRequest is the request for transaction data retrieval
// @Param UserID int true "ID of the user"
// @Param CartID int true "ID of the cart"
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
