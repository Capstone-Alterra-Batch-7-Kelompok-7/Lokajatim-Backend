package response

import (
	"lokajatim/entities"
	"time"
)

// ProductResponse is the response for the Product controller
// @Description ProductResponse is the response for the Product controller
// @Param ID int true "ID of the product"
// @Param Name string true "Name of the product"
// @Param Price int true "Price of the product"
// @Param Stock int true "Stock of the product"
// @Param Description string true "Description of the product"
// @Param Photo string true "URL Photo of the product"
// @Param CategoryID int true "Category ID of the product"
// @Param Category entities.Category true "Category of the product"
// @Param CreatedAt string true "Created At of the product"
// @Param UpdatedAt string true "Updated At of the product"
type ProductResponse struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Price       int               `json:"price"`
	Stock       int               `json:"stock"`
	Description string            `json:"description"`
	Photo       string            `json:"photo"`
	CategoryID  int               `json:"category_id"`
	Category    entities.Category `json:"Category"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

func ProductFromEntities(product entities.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.Description,
		Photo:       product.Photo,
		CategoryID:  product.CategoryID,
		Category:    product.Category,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
