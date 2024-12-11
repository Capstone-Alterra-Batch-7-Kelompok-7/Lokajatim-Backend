package payment

import (
	"lokajatim/entities"
	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepoInterface {
	return &PaymentRepositoryImpl{db: db}
}

func (r *PaymentRepositoryImpl) CreatePayment(payment entities.Payment) (entities.Payment, error) {
	if err := r.db.Create(&payment).Error; err != nil {
		return entities.Payment{}, err
	}

	var createdPayment entities.Payment
	result := r.db.Preload("Transaction").First(&createdPayment, payment.ID)
	if result.Error != nil {
		return entities.Payment{}, result.Error
	}
	return createdPayment, nil
}

func (r *PaymentRepositoryImpl) GetPaymentByID(paymentID int) (entities.Payment, error) {
	var payment entities.Payment
	result := r.db.Preload("Transaction").First(&payment, paymentID)
	if result.Error != nil {
		return entities.Payment{}, result.Error
	}
	return payment, nil
}

func (r *PaymentRepositoryImpl) UpdatePaymentStatus(paymentID int, status string) error {
	if err := r.db.Model(&entities.Payment{}).Where("id = ?", paymentID).Update("payment_status", status).Error; err != nil {
		return err
	}
	return nil
}

func (r *PaymentRepositoryImpl) DeletePayment(paymentID int) error {
	if err := r.db.Where("id = ?", paymentID).Delete(&entities.Payment{}).Error; err != nil {
		return err
	}
	return nil
}