package payment

import "lokajatim/entities"

type PaymentRepoInterface interface {
	CreatePayment(payment entities.Payment) (entities.Payment, error)
	GetPaymentByID(paymentID int) (entities.Payment, error)
	UpdatePaymentStatus(paymentID int, status string) error
	DeletePayment(paymentID int) error
}
