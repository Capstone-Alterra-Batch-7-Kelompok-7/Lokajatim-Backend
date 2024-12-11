package transaction

import (
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
	if err := r.db.Create(&transaction).Error; err != nil {
		return entities.Transaction{}, err
	}

	var createdTransaction entities.Transaction
	result := r.db.Preload("User").Preload("Cart").First(&createdTransaction, transaction.ID)
	if result.Error != nil {
		return entities.Transaction{}, result.Error
	}
	return createdTransaction, nil
}

func (r *TransactionRepositoryImpl) GetTransactionByID(transactionID int) (entities.Transaction, error) {
	var transaction entities.Transaction
	result := r.db.Preload("User").Preload("Cart").First(&transaction, transactionID)
	if result.Error != nil {
		return entities.Transaction{}, result.Error
	}
	return transaction, nil
}

func (r *TransactionRepositoryImpl) GetAllTransactions() ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	result := r.db.Preload("User").Preload("Cart").Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func (r *TransactionRepositoryImpl) UpdateTransaction(transactionID int, transaction entities.Transaction) (entities.Transaction, error) {
	if err := r.db.Model(&entities.Transaction{}).Where("id = ?", transactionID).Updates(transaction).Error; err != nil {
		return entities.Transaction{}, err
	}

	var updatedTransaction entities.Transaction
	result := r.db.Preload("User").Preload("Cart").First(&updatedTransaction, transactionID)
	if result.Error != nil {
		return entities.Transaction{}, result.Error
	}
	return updatedTransaction, nil
}

func (r *TransactionRepositoryImpl) UpdateTransactionStatus(transactionID int, status string) error {
	if err := r.db.Model(&entities.Transaction{}).Where("id = ?", transactionID).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepositoryImpl) DeleteTransaction(transactionID int) error {
	if err := r.db.Model(&entities.Transaction{}).Where("id = ?", transactionID).Delete(&entities.Transaction{}).Error; err != nil {
		return err
	}
	return nil
}