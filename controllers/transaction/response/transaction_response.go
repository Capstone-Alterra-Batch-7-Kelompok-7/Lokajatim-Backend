package response

import (
	"lokajatim/entities"
	"time"
)

// TransactionResponse is the response for the transaction controller
// @Description TransactionResponse is the response for transaction data retrieval
// @Param ID int true "ID of the transaction"
// @Param TransactionID string true "Transaction ID of the transaction"
// @Param User UserResponse true "User of the transaction"
// @Param CartID int true "Cart ID of the transaction"
// @Param TotalPrice float64 true "Total price of the transaction"
// @Param Status string true "Status of the transaction"
// @Param Products []ProductInCart true "Products in the cart"
// @Param PaymentURL string true "Payment URL of the transaction"
// @Param CreatedAt string true "Created at of the transaction"
// @Param UpdatedAt string true "Updated at of the transaction"
type TransactionResponse struct {
	ID            int             `json:"id"`
	TransactionID string          `json:"transaction_id"`
	User          UserResponse    `json:"user"`
	CartID        int             `json:"cart_id"`
	TotalPrice    float64         `json:"total_price"`
	Status        string          `json:"status"`
	PaymentURL    string          `json:"payment_url"`
	Products      []ProductInCart `json:"products"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

// ProductInCart is the response for the product in cart
// @Description ProductInCart is the response for product in cart data retrieval
// @Param Name string true "Name of the product"
// @Param Category string true "Category of the product"
// @Param Quantity int true "Quantity of the product"
// @Param TotalPrice float64 true "Total price of the product"
type ProductInCart struct {
	Name       string `json:"name"`
	Category   string `json:"category"`
	Quantity   int    `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

// UserResponse is the response for the user controller
// @Description UserResponse is the response for user data retrieval
// @Param ID int true "ID of the user"
// @Param Name string true "Name of the user"
// @Param Email string true "Email of the user"
// @Param Address string true "Address of the user"
// @Param PhoneNumber string true "Phone number of the user"
type UserResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

func TransactionFromEntity(transaction entities.Transaction) TransactionResponse {

	userResponse := UserResponse{
		ID:          transaction.User.ID,
		Name:        transaction.User.Name,
		Email:       transaction.User.Email,
		Address:     transaction.User.Address,
		PhoneNumber: transaction.User.PhoneNumber,
	}

	var products []ProductInCart
	for _, item := range transaction.Cart.Items {
		product := ProductInCart{
			Name:       item.Product.Name,
			Category:   item.Product.Category.Name,
			Quantity:   item.Quantity,
			TotalPrice: float64(item.Product.Price) * float64(item.Quantity),
		}
		products = append(products, product)
	}

	return TransactionResponse{
		ID:            transaction.ID,
		TransactionID: transaction.TransactionID,
		User:          userResponse,
		CartID:        transaction.CartID,
		TotalPrice:    transaction.TotalPrice,
		Status:        transaction.Status,
		Products:      products,
		PaymentURL:    transaction.PaymentURL,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}
