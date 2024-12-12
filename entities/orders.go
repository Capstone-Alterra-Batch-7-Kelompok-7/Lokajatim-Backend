package entities

import (
	"time"
)

type Order struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	UserID      int       `json:"user_id"`
	User        User      `json:"user"`
	ProductID   int       `json:"product_id"`
	Product     Product   `json:"product"`
	TotalAmount int       `json:"total_amount"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	NIK         string    `json:"nik"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
