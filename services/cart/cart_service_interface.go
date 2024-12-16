package cart

import "lokajatim/entities"

type CartServiceInterface interface {
	FindByUserID(userID int) (entities.Cart, error)
	FindByID(cartID int) (entities.Cart, error)
	FindByCartItemID(cartItemID int) (entities.Cart, error)
	Create(cart entities.Cart) (entities.Cart, error)
	AddItemToCart(userID int, cartItem entities.CartItem) (entities.CartItem, error)
	UpdateItemQuantity(cartItemID int, quantity int) (entities.CartItem, error)
	RemoveItemFromCart(cartItemID int) error
	ClearCart(cartID int) error
}