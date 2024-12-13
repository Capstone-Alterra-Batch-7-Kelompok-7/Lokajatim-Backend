package transaction

import "lokajatim/entities"

type TransactionRepositoryInterface interface {
	CreateTransaction(transaction entities.Transaction) (entities.Transaction, error)
	GetTransactionByID(transactionID int) (entities.Transaction, error)
	GetAllTransactions() ([]entities.Transaction, error)
	UpdateTransaction(transactionID int, updates map[string]interface{}) (entities.Transaction, error)
	UpdateTransactionStatus(transactionID int, status string) error
	DeleteTransaction(transactionID int) error
}
