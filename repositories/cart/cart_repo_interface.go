package cart

import "lokajatim/entities"

type CartRepositoryInterface interface {
	GetCartbyUserID(userID int) (entities.Cart, error)
	AddItemToCart(cartItem entities.CartItem) (entities.CartItem, error)
	UpdateItemQuantity(cartItemID, quantity int) error
	RemoveItemFromCart(cartItemID int) error
	ClearCart(cartID int) error
}