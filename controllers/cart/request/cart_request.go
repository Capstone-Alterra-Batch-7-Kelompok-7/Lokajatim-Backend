package request

import "lokajatim/entities"

type CartRequest struct {
	UserID    int               `json:"user_id" validate:"required"`
	CartItems []CartItemRequest `json:"cart_items" validate:"required"`
}

type CartItemRequest struct {
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type QuantityRequest struct {
	Quantity int `json:"quantity" validate:"required"`
}

func (cartRequest CartRequest) ToEntities() (entities.Cart, []entities.CartItem, error) {
	cart := entities.Cart{
		UserID: cartRequest.UserID,
	}

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

func (cartItemRequest CartItemRequest) ToEntities() (entities.CartItem, error) {
	cartItem := entities.CartItem{
		ProductID: cartItemRequest.ProductID,
		Quantity:  cartItemRequest.Quantity,
	}
	return cartItem, nil
}

func (quantityRequest QuantityRequest) ToEntities() (entities.CartItem, error) {
	cartItem := entities.CartItem{
		Quantity: quantityRequest.Quantity,
	}
	return cartItem, nil
}
