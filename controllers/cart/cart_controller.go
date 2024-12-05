package cart

import (
	"lokajatim/controllers/base"
	"lokajatim/controllers/cart/request"
	"lokajatim/controllers/cart/response"
	"lokajatim/services/cart"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartController struct {
	CartService cart.CartService
}

func NewCartController(service cart.CartService) *CartController {
	return &CartController{CartService: service}
}

func (h *CartController) GetCartByUserID(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	cart, err := h.CartService.GetCartbyUserID(userID)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to get cart",
		})
	}
	return base.SuccesResponse(c, response.CartFromEntities(cart))
}

func (h *CartController) AddItemToCart(c echo.Context) error {
	req := new(request.CartItemRequest)
	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to bind request",
		})
	}

	cartItem, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to convert request to entities",
		})
	}

	// Retrieve user ID from context
	userID := c.Get("user_id").(int) // Assumes user_id is stored in the context

	// Add the item to the cart using the service, handle both return values (CartItem, error)
	_, err = h.CartService.AddItemToCart(cartItem) 
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to add item to cart",
		})
	}

	// Fetch the updated cart for the user
	updatedCart, err := h.CartService.GetCartbyUserID(userID)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to fetch updated cart",
		})
	}

	// Return success response with updated cart
	return base.SuccesResponse(c, response.CartFromEntities(updatedCart))
}


func (h *CartController) UpdateItemQuantity(c echo.Context) error {
	cartItemID, _ := strconv.Atoi(c.Param("cart_item_id"))
	var quantity struct {
		Quantity int `json:"quantity" validate:"required"`
	}
	if err := c.Bind(&quantity); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to bind quantity",
		})
	}

	err := h.CartService.UpdateItemQuantity(cartItemID, quantity.Quantity)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to add item quantity",
		})
	}

	return base.SuccesResponse(c, "Item quantity added successfully")
}

func (h *CartController) RemoveItemFromCart(c echo.Context) error {
	cartItemID, _ := strconv.Atoi(c.Param("cart_item_id"))
	err := h.CartService.RemoveItemFromCart(cartItemID)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to remove item from cart",
		})
	}
	return base.SuccesResponse(c, "Item removed from cart successfully")
}

func (h *CartController) ClearCart(c echo.Context) error {
	cartID, _ := strconv.Atoi(c.Param("cart_id"))
	err := h.CartService.ClearCart(cartID)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to clear cart",
		})
	}
	return base.SuccesResponse(c, "Cart cleared successfully")
}
