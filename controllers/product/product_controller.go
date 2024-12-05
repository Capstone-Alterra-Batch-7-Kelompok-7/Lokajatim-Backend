package product

import (
	"lokajatim/controllers/base"
	"lokajatim/controllers/pagination"
	"lokajatim/controllers/product/request"
	"lokajatim/controllers/product/response"
	"lokajatim/services/product"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	ProductService product.ProductService
}

func NewProductController(service product.ProductService) *ProductController {
	return &ProductController{ProductService: service}
}

// @Summary Get all products
// @Description Get all products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} response.ProductResponse
// @Failure 400 {object} base.BaseResponse
// @Router /products [get]
func (h *ProductController) GetAllProducts(c echo.Context) error {
	products, err := h.ProductService.GetProducts()
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to get products",
		})
	}
	return pagination.SuccessPaginatedResponse(c, products, 1, 10, int64(len(products)))
}

// @Summary Get product by ID
// @Description Get product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "ID of the product"
// @Success 200 {object} response.ProductResponse
// @Failure 400 {object} base.BaseResponse
// @Router /products/{id} [get]
func (h *ProductController) GetProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.ProductService.GetProductByID(id)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to get product",
		})
	}
	photos, err := h.ProductService.GetProductPhotos(id)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to get product photos",
		})
	}
	return base.SuccesResponse(c, response.ProductFromEntities(product, photos))
}

// @Summary Create product
// @Description Create product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body request.ProductRequest true "Product body"
// @Success 201 {object} response.ProductResponse
// @Failure 400 {object} base.BaseResponse
// @Router /products [post]
func (h *ProductController) CreateProduct(c echo.Context) error {
	req := new(request.ProductRequest)
	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to bind request",
		})
	}

	// Convert to entity
	product, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to convert request to entities",
		})
	}

	// Create the product
	createdProduct, err := h.ProductService.CreateProduct(product)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to create product",
		})
	}

	// Create product photos
	photos, err := req.ToProductPhotos(createdProduct.ID)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to convert photos to entities",
		})
	}

	// Save photos
	err = h.ProductService.CreateProductPhotos(photos)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to save product photos",
		})
	}

	return base.SuccesResponse(c, response.ProductFromEntities(createdProduct, photos))
}

// @Summary Update product
// @Description Update product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "ID of the product"
// @Param product body request.ProductRequest true "Product body"
// @Success 200 {object} response.ProductResponse
// @Failure 400 {object} base.BaseResponse
// @Router /products/{id} [put]
func (h *ProductController) UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := new(request.ProductRequest)
	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to bind request",
		})
	}

	// Convert to entity
	product, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to convert request to entities",
		})
	}

	// Update the product
	updatedProduct, err := h.ProductService.UpdateProduct(id, product)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to update product",
		})
	}

	// Handle product photos
	photos, err := req.ToProductPhotos(updatedProduct.ID)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to convert photos to entities",
		})
	}

	// Delete old photos before adding new ones (if applicable)
	err = h.ProductService.UpdateProductPhotos(updatedProduct.ID, photos)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to update product photos",
		})
	}

	return base.SuccesResponse(c, response.ProductFromEntities(updatedProduct, photos))
}

// @Summary Delete product
// @Description Delete product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "ID of the product"
// @Success 200 "Product deleted successfully"
// @Failure 400 {object} base.BaseResponse
// @Router /products/{id} [delete]
func (h *ProductController) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.ProductService.DeleteProduct(id)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to delete product",
		})
	}
	return base.SuccesResponse(c, "Product deleted successfully")
}