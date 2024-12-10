package cart

import (
	"lokajatim/entities"

	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepositoryInterface {
	return &cartRepository{db: db}
}

func (r *cartRepository) FindByUserID(userID int) (entities.Cart, error) {
	var cart entities.Cart
	result := r.db.Preload("Items.Product").Preload("User").First(&cart, "user_id = ?", userID)
	if result.Error != nil {
		return entities.Cart{}, result.Error
	}
	return cart, nil
}

// FindByID finds a cart by its ID
func (r *cartRepository) FindByID(cartID int) (entities.Cart, error) {
	var cart entities.Cart
	result := r.db.Preload("Items.Product").Preload("User").First(&cart, cartID)
	if result.Error != nil {
		return entities.Cart{}, result.Error
	}
	return cart, nil
}

// Create creates a new cart in the database
func (r *cartRepository) Create(cart entities.Cart) (entities.Cart, error) {
	if err := r.db.Create(&cart).Error; err != nil {
		return entities.Cart{}, err
	}
	return cart, nil
}

// AddItemToCart adds a CartItem to the Cart
func (r *cartRepository) AddItemToCart(userID int, cartItem entities.CartItem) (entities.CartItem, error) {
	var cart entities.Cart
	result := r.db.Preload("Items.Product").Preload("User").First(&cart, "user_id = ?", userID)

	if result.Error != nil {
		cart = entities.Cart{UserID: userID}
		if err := r.db.Create(&cart).Error; err != nil {
			return entities.CartItem{}, err
		}
	}

	cartItem.CartID = cart.ID
	if err := r.db.Create(&cartItem).Error; err != nil {
		return entities.CartItem{}, err
	}
	cart.CalculateTotalPrice()
	if err := r.db.Save(&cart).Error; err != nil {
		return entities.CartItem{}, err
	}

	return cartItem, nil
}

// UpdateItemQuantity updates the quantity of a CartItem
func (r *cartRepository) UpdateItemQuantity(cartItemID int, quantity int) (entities.CartItem, error) {
	var cartItem entities.CartItem
	if err := r.db.Preload("Product").First(&cartItem, cartItemID).Error; err != nil {
		return entities.CartItem{}, err
	}
	cartItem.Quantity = quantity
	if err := r.db.Save(&cartItem).Error; err != nil {
		return entities.CartItem{}, err
	}
	return cartItem, nil
}

// RemoveItemFromCart removes a CartItem from the Cart
func (r *cartRepository) RemoveItemFromCart(cartItemID int) error {
	var cartItem entities.CartItem
	if err := r.db.First(&cartItem, cartItemID).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&cartItem).Error; err != nil {
		return err
	}
	return nil
}

// ClearCart removes all items in the Cart
func (r *cartRepository) ClearCart(cartID int) error {
	var cart entities.Cart

	if err := r.db.First(&cart, cartID).Error; err != nil {
		return err
	}

	if err := r.db.Where("cart_id = ?", cartID).Delete(&entities.CartItem{}).Error; err != nil {
		return err
	}

	cart.CalculateTotalPrice()
	if err := r.db.Save(&cart).Error; err != nil {
		return err
	}

	return nil
}
