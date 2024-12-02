package entities

import (
	"time"
)

type Product struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	Photo       string    `json:"photo"`
	CategoryID  int       `json:"category_id"`
	Category    Category  `json:"Category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
