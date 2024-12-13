package entities

import (
	"time"
)

type Product struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Price       int            `json:"price"`
	Stock       int            `json:"stock"`
	Description string         `json:"description"`
	CategoryID  int            `json:"category_id"`
	Category    Category       `json:"category"`
	Photos      []ProductPhoto `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;" json:"photos"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
