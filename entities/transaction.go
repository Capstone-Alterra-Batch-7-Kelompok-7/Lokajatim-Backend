package entities

import "time"

type Transaction struct {
	ID            int       `gorm:"primary_key;auto_increment" json:"id"`
	TransactionID string    `gorm:"unique_index" json:"transaction_id"`
	UserID        int       `json:"user_id"`
	User          User      `json:"user"`
	CartID        int       `json:"cart_id"`
	Cart          Cart      `json:"cart"`
	TotalPrice    float64   `json:"total_price"`
	Status        string    `json:"status"`
	PaymentURL    string    `json:"payment_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
