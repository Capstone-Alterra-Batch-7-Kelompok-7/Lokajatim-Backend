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

func (s *CartService) FindByUserID(userID int) (entities.Cart, error) {
	return s.cartRepository.FindByUserID(userID)
}

func (s *CartService) FindByID(cartID int) (entities.Cart, error) {
	return s.cartRepository.FindByID(cartID)
}

func (s *CartService) FindByCartItemID(cartItemID int) (entities.Cart, error) {
	return s.cartRepository.FindByCartItemID(cartItemID)
}

func (s *CartService) Create(cart entities.Cart) (entities.Cart, error) {
	return s.cartRepository.Create(cart)
}

func (s *CartService) AddItemToCart(userID int, cartItem entities.CartItem) (entities.CartItem, error) {
	return s.cartRepository.AddItemToCart(userID, cartItem)
}

func (s *CartService) UpdateItemQuantity(cartItemID int, quantity int) (entities.CartItem, error) {
	return s.cartRepository.UpdateItemQuantity(cartItemID, quantity)
}

func (s *CartService) RemoveItemFromCart(cartItemID int) error {
	return s.cartRepository.RemoveItemFromCart(cartItemID)
}

func (s *CartService) ClearCart(cartID int) error {
	return s.cartRepository.ClearCart(cartID)
}