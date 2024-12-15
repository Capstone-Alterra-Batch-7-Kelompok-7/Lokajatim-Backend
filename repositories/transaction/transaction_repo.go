package transaction

import (
	"errors"
	"lokajatim/entities"

	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepositoryInterface {
	return &TransactionRepositoryImpl{db: db}
}

func (r *TransactionRepositoryImpl) CreateTransaction(transaction entities.Transaction) (entities.Transaction, error) {
	if transaction.UserID == 0 || transaction.CartID == 0 {
		return entities.Transaction{}, errors.New("user_id or cart_id is missing")
	}
	if err := r.db.Create(&transaction).Error; err != nil {
		return entities.Transaction{}, err
	}

	var createdTransaction entities.Transaction
	if err := r.db.Preload("User").First(&createdTransaction, transaction.ID).Error; err != nil {
		return entities.Transaction{}, err
	}
	return createdTransaction, nil
}

func (r *TransactionRepositoryImpl) GetTransactionByID(transactionID int) (entities.Transaction, error) {
	var transaction entities.Transaction
	if err := r.db.Preload("User").First(&transaction, transactionID).Error; err != nil {
		return entities.Transaction{}, err
	}
	return transaction, nil
}

func (r *TransactionRepositoryImpl) GetAllTransactions() ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	if err := r.db.Preload("User").Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepositoryImpl) UpdateTransaction(transactionID int, updates map[string]interface{}) (entities.Transaction, error) {
	if err := r.db.Model(&entities.Transaction{}).Where("id = ?", transactionID).Updates(updates).Error; err != nil {
		return entities.Transaction{}, err
	}

	var updatedTransaction entities.Transaction
	if err := r.db.Preload("User").First(&updatedTransaction, transactionID).Error; err != nil {
		return entities.Transaction{}, err
	}
	return updatedTransaction, nil
}

func (r *TransactionRepositoryImpl) UpdateTransactionStatus(transactionID int, status string) error {
	return r.db.Model(&entities.Transaction{}).Where("id = ?", transactionID).Update("status", status).Error
}

func (r *TransactionRepositoryImpl) DeleteTransaction(transactionID int) error {
	result := r.db.Where("id = ?", transactionID).Delete(&entities.Transaction{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("transaction not found")
	}
	return nil
}

func (r *TransactionRepositoryImpl) GetTransactionByOrderID(orderID string) (entities.Transaction, error) {
	var transaction entities.Transaction
	if err := r.db.Preload("User").Where("transaction_id = ?", orderID).First(&transaction).Error; err != nil {
		return entities.Transaction{}, err
	}
	return transaction, nil
}

func (r *TransactionRepositoryImpl) UpdateTransactionStatusByOrderID(orderID string, status string) error {
	return r.db.Model(&entities.Transaction{}).Where("transaction_id = ?", orderID).Update("status", status).Error
}
