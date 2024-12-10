package entities

import (
	"time"
)

type Cart struct {
	ID                         int        `json:"id" gorm:"primaryKey"`
	UserID                     int        `json:"user_id"`
	User                       User       `json:"user"`
	Items                      []CartItem `json:"items" gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE;"`
	TotalPrice                 float64    `json:"total_price" gorm:"-"`
	TotalPriceAfterTransaction float64    `json:"total_price_after_transaction" gorm:"-"`
	CreatedAt                  time.Time  `json:"created_at"`
	UpdatedAt                  time.Time  `json:"updated_at"`
}

type CartItem struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	CartID    int       `json:"cart_id"`
	Cart      Cart      `json:"cart"`
	ProductID int       `json:"product_id"`
	Product   Product   `json:"product"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (cart *Cart) CalculateTotalPrice() {
	total := 0.0
	for _, item := range cart.Items {
		total += float64(item.Product.Price) * float64(item.Quantity)
	}
	cart.TotalPrice = total
}

func (cart *Cart) CalculateTotalPriceAfterAddTransactionPrice() {
	cart.CalculateTotalPrice()
	cart.TotalPriceAfterTransaction = cart.TotalPrice + 2000
}
