package cart

import (
	"lokajatim/entities"

	"gorm.io/gorm"
)

type CartRepositoryImpl struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepositoryInterface {
	return &CartRepositoryImpl{db: db}
}

func (r *CartRepositoryImpl) GetCartbyUserID(userID int) (entities.Cart, error) {
	var cart entities.Cart
	err := r.db.Preload("Items.Product").Where("user_id = ?", userID).Find(&cart).Error
	if err != nil {
		return entities.Cart{}, err
	}
	cart.CalculateTotalPrice()

	return cart, nil
}

func (r *CartRepositoryImpl) AddItemToCart(cartItem entities.CartItem) (entities.CartItem, error) {
	var existingItem entities.CartItem
	err := r.db.Where("cart_id = ? AND product_id = ?", cartItem.CartID, cartItem.ProductID).First(&existingItem).Error
	if err == nil {
		existingItem.Quantity += cartItem.Quantity
		if updateErr := r.db.Save(&existingItem).Error; updateErr != nil {
			return entities.CartItem{}, updateErr
		}
		return existingItem, nil
	}

	if err := r.db.Create(&cartItem).Error; err != nil {
		return entities.CartItem{}, err
	}
	return cartItem, nil
}

func (r *CartRepositoryImpl) UpdateItemQuantity(cartItemID, quantity int) error {
	var cartItem entities.CartItem
	if err := r.db.First(&cartItem, cartItemID).Error; err != nil {
		return err
	}
	cartItem.Quantity += quantity

	return r.db.Model(&cartItem).Update("quantity", quantity).Error
}

func (r *CartRepositoryImpl) RemoveItemFromCart(cartItemID int) error {
	return r.db.Delete(&entities.CartItem{}, cartItemID).Error
}

func (r *CartRepositoryImpl) ClearCart(cartID int) error {
	return r.db.Where("cart_id = ?", cartID).Delete(&entities.CartItem{}).Error
}
