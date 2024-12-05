package cart

import (
	"lokajatim/entities"
	"lokajatim/repositories/cart"
)

type CartService struct {
	cartRepository cart.CartRepositoryInterface
}

func NewCartService(cartRepo cart.CartRepositoryInterface) *CartService {
	return &CartService{cartRepository: cartRepo}
}

func (s *CartService) GetCartbyUserID(userID int) (entities.Cart, error) {
	return s.cartRepository.GetCartbyUserID(userID)
}

func (s *CartService) AddItemToCart(cartItem entities.CartItem) (entities.CartItem, error) {
	return s.cartRepository.AddItemToCart(cartItem)
}

func (s *CartService) UpdateItemQuantity(cartItemID, quantity int) error {
	return s.cartRepository.UpdateItemQuantity(cartItemID, quantity)
}

func (s *CartService) RemoveItemFromCart(cartItemID int) error {
	return s.cartRepository.RemoveItemFromCart(cartItemID)
}

func (s *CartService) ClearCart(cartID int) error {
	return s.cartRepository.ClearCart(cartID)
}
