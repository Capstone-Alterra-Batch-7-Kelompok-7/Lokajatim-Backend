package entities

import "time"

type Ticket struct {
	ID           uint      `json:"id"`
	EventsID     uint      `json:"events_id"`
	Quantity     int       `json:"quantity"`
	TotalAmount  int       `json:"total_amount"`
	UsersID      int      `json:"users_id"`
	PaymentMethod string   `json:"payment_method"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
