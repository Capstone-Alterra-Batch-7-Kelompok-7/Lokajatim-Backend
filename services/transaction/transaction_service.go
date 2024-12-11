package transaction

import (
	"lokajatim/entities"
	"lokajatim/repositories/transaction"
)

type TransactionService struct {
	transactionRepo transaction.TransactionRepositoryInterface
}

func NewTransactionService(transactionRepo transaction.TransactionRepositoryInterface) TransactionService {
	return TransactionService{transactionRepo: transactionRepo}
}

func (s *TransactionService) CreateTransaction(transaction entities.Transaction) (entities.Transaction, error) {
	return s.transactionRepo.CreateTransaction(transaction)
}

func (s *TransactionService) GetTransactionByID(transactionID int) (entities.Transaction, error) {
	return s.transactionRepo.GetTransactionByID(transactionID)
}

func (s *TransactionService) GetAllTransactions() ([]entities.Transaction, error) {
	return s.transactionRepo.GetAllTransactions()
}

func (s *TransactionService) UpdateTransaction(transactionID int, transaction entities.Transaction) (entities.Transaction, error) {
	return s.transactionRepo.UpdateTransaction(transactionID, transaction)
}

func (s *TransactionService) UpdateTransactionStatus(transactionID int, status string) error {
	return s.transactionRepo.UpdateTransactionStatus(transactionID, status)
}

func (s *TransactionService) DeleteTransaction(transactionID int) error {
	return s.transactionRepo.DeleteTransaction(transactionID)
}