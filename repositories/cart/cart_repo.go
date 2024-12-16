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
	result := r.db.Preload("Items.Product.Photos").
        Preload("Items.Product").
        Preload("Items").
        Preload("User").First(&cart, "user_id = ?", userID)
	if result.Error != nil {
		return entities.Cart{}, result.Error
	}
	cart.CalculateTotalPrice()
	cart.CalculateTotalPriceAfterAddTransactionPrice()
	if err := r.db.Save(&cart).Error; err != nil {
		return entities.Cart{}, err
	}
	return cart, nil
}

func (r *cartRepository) FindByID(cartID int) (entities.Cart, error) {
	var cart entities.Cart
	result := r.db.Preload("Items.Product.Photos").
        Preload("Items.Product").
        Preload("Items").
        Preload("User").First(&cart, cartID)
	if result.Error != nil {
		return entities.Cart{}, result.Error
	}
	cart.CalculateTotalPrice()
	cart.CalculateTotalPriceAfterAddTransactionPrice()
	if err := r.db.Save(&cart).Error; err != nil {
		return entities.Cart{}, err
	}
	return cart, nil
}

func (r *cartRepository) FindByCartItemID(cartItemID int) (entities.Cart, error) {
    var cartItem entities.CartItem
    result := r.db.Preload("Cart.Items.Product.Photos").Preload("Cart.Items.Product").
        Preload("Cart.User").First(&cartItem, cartItemID)
    if result.Error != nil {
        return entities.Cart{}, result.Error
    }

    var cart entities.Cart
    result = r.db.Preload("Items.Product.Photos").
        Preload("Items.Product").
        Preload("Items").
        Preload("User").
        First(&cart, cartItem.CartID)
    if result.Error != nil {
        return entities.Cart{}, result.Error
    }

    cart.CalculateTotalPrice()
    cart.CalculateTotalPriceAfterAddTransactionPrice()
    if err := r.db.Save(&cart).Error; err != nil {
        return entities.Cart{}, err
    }

    return cart, nil
}

func (r *cartRepository) Create(cart entities.Cart) (entities.Cart, error) {
	if err := r.db.Create(&cart).Error; err != nil {
		return entities.Cart{}, err
	}
	return cart, nil
}

func (r *cartRepository) AddItemToCart(userID int, cartItem entities.CartItem) (entities.CartItem, error) {
	var cart entities.Cart
	result := r.db.Preload("Items.Product.Photos").
        Preload("Items.Product").
        Preload("Items").
        Preload("User").First(&cart, "user_id = ?", userID)

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
	cart.CalculateTotalPriceAfterAddTransactionPrice()
	if err := r.db.Save(&cart).Error; err != nil {
		return entities.CartItem{}, err
	}

	return cartItem, nil
}

func (r *cartRepository) UpdateItemQuantity(cartItemID int, quantity int) (entities.CartItem, error) {
	var cartItem entities.CartItem
	if err := r.db.Preload("Product").First(&cartItem, cartItemID).Error; err != nil {
		return entities.CartItem{}, err
	}
	cartItem.Quantity = quantity
	if err := r.db.Save(&cartItem).Error; err != nil {
		return entities.CartItem{}, err
	}

	var cart entities.Cart
	if err := r.db.Preload("Items.Product").First(&cart, cartItem.CartID).Error; err == nil {
		cart.CalculateTotalPrice()
		cart.CalculateTotalPriceAfterAddTransactionPrice()
		_ = r.db.Save(&cart)
	}

	return cartItem, nil
}

func (r *cartRepository) RemoveItemFromCart(cartItemID int) error {
	var cartItem entities.CartItem
	if err := r.db.First(&cartItem, cartItemID).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&cartItem).Error; err != nil {
		return err
	}

	var cart entities.Cart
	if err := r.db.Preload("Items.Product").First(&cart, cartItem.CartID).Error; err == nil {
		cart.CalculateTotalPrice()
		cart.CalculateTotalPriceAfterAddTransactionPrice()
		_ = r.db.Save(&cart)
	}

	return nil
}

func (r *cartRepository) ClearCart(cartID int) error {
	var cart entities.Cart
	if err := r.db.First(&cart, cartID).Error; err != nil {
		return err
	}

	if err := r.db.Where("cart_id = ?", cartID).Delete(&entities.CartItem{}).Error; err != nil {
		return err
	}

	cart.CalculateTotalPrice()
	cart.CalculateTotalPriceAfterAddTransactionPrice()
	if err := r.db.Save(&cart).Error; err != nil {
		return err
	}

	return nil
}