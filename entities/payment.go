package entities

import (
	"time"
)

type Payment struct {
	ID            int         `gorm:"primary_key;auto_increment" json:"id"`
	TransactionID string      `json:"transaction_id"`
	Transaction   Transaction `json:"transaction"`
	PaymentMethod string      `json:"payment_method"`
	PaymentStatus string      `json:"payment_status"`
	PaymentURL    string      `json:"payment_url"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}
