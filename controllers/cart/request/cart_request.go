package request

import "lokajatim/entities"

// CartRequest is the request for creating a Cart and adding CartItems at the same time
// @Description CartRequest is the request for creating a Cart and adding CartItems in one request
// @Param user_id body int true "User ID"
// @Param cart_items body array true "List of Cart Items to be added"
type CartRequest struct {
	UserID    int               `json:"user_id" validate:"required"`
	CartItems []CartItemRequest `json:"cart_items" validate:"required"`
}

// CartItemRequest is the request for the CartItem endpoint
// @Description CartItemRequest is the request for adding products to the Cart
// @Param product_id body int true "Product ID"
// @Param quantity body int true "Quantity of the product"
type CartItemRequest struct {
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

// ToEntities converts CartRequest to Cart entity, also converts CartItems to CartItem entities
func (cartRequest CartRequest) ToEntities() (entities.Cart, []entities.CartItem, error) {
	// Create Cart entity
	cart := entities.Cart{
		UserID: cartRequest.UserID,
	}

	// Convert CartItems from CartItemRequest to CartItem entities
	var cartItems []entities.CartItem
	for _, itemRequest := range cartRequest.CartItems {
		cartItem, err := itemRequest.ToEntities()
		if err != nil {
			return cart, nil, err
		}
		cartItems = append(cartItems, cartItem)
	}

	return cart, cartItems, nil
}

// ToEntities converts CartItemRequest to CartItem entity
func (cartItemRequest CartItemRequest) ToEntities() (entities.CartItem, error) {
	cartItem := entities.CartItem{
		ProductID: cartItemRequest.ProductID,
		Quantity:  cartItemRequest.Quantity,
	}
	return cartItem, nil
}
