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

// @Summary Get cart by user ID
// @Description Get cart by user ID
// @Tags Cart
// @Accept json
// @Produce json
// @Param user_id path int true "ID of the user"
// @Success 200 {object} response.CartResponse
// @Failure 400 {object} base.BaseResponse
// @Router /carts/{user_id} [get]
func (h *CartController) GetCartByUserID(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	cart, err := h.CartService.FindByUserID(userID)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to get cart",
		})
	}
	if cart.ID == 0 {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Cart not found",
		})
	}
	return base.SuccesResponse(c, response.CartFromEntities(cart))
}

// @Summary Add item to cart
// @Description Add item to cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param user_id path int true "ID of the user"
// @Param request body request.CartRequest true "Cart Request"
// @Success 200 {object} response.CartResponse
// @Failure 400 {object} base.BaseResponse
// @Router /carts [post]
func (h *CartController) AddItemToCart(c echo.Context) error {
	req := new(request.CartRequest)
	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to bind request",
		})
	}

	cart, cartItems, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to convert request to entities",
		})
	}

	for _, cartItem := range cartItems {
		_, err = h.CartService.AddItemToCart(cart.UserID, cartItem)
		if err != nil {
			return base.ErrorResponse(c, err, map[string]string{
				"error": "Failed to add item to cart",
			})
		}
	}

	cart, err = h.CartService.FindByUserID(cart.UserID)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to fetch updated cart",
		})
	}
	return base.SuccesResponse(c, response.CartFromEntities(cart))
}

// @Summary Update item quantity
// @Description Update item quantity
// @Tags Cart
// @Accept json
// @Produce json
// @Param cart_item_id path int true "ID of the cart item"
// @Param request body request.QuantityRequest true "Quantity Request"
// @Success 200 {object} response.CartItemResponse
// @Failure 400 {object} base.BaseResponse
// @Router /carts/{cart_item_id} [put]
func (h *CartController) UpdateItemQuantity(c echo.Context) error {
	cartItemID, _ := strconv.Atoi(c.Param("cart_item_id"))
	var quantity request.QuantityRequest
	if err := c.Bind(&quantity); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to bind quantity",
		})
	}

	cartItem, err := h.CartService.UpdateItemQuantity(cartItemID, quantity.Quantity)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to update item quantity",
		})
	}

	return base.SuccesResponse(c, response.CartItemResponse{
		ID:        cartItem.ID,
		ProductID: cartItem.ProductID,
		Product: response.ProductResponse{
			ID:          cartItem.Product.ID,
			Name:        cartItem.Product.Name,
			Price:       cartItem.Product.Price,
			Stock:       cartItem.Product.Stock,
			Description: cartItem.Product.Description,
			CreatedAt:   cartItem.Product.CreatedAt,
			UpdatedAt:   cartItem.Product.UpdatedAt,
		},
		Quantity:  cartItem.Quantity,
		CreatedAt: cartItem.CreatedAt,
		UpdatedAt: cartItem.UpdatedAt,
	})
}

// @Summary Remove item from cart
// @Description Remove item from cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param cart_item_id path int true "ID of the cart item"
// @Success 200 {object} base.BaseResponse
// @Failure 400 {object} base.BaseResponse
// @Router /carts/{cart_item_id} [delete]
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

// @Summary Clear cart
// @Description Clear cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param cart_id path int true "ID of the cart"
// @Success 200 {object} base.BaseResponse
// @Failure 400 {object} base.BaseResponse
// @Router /carts/{cart_id}/clear [delete]
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
