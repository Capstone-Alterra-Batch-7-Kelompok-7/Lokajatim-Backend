package response

import (
	"lokajatim/entities"
	"time"
)

// ProductResponse is the response for the product controller
// @Description ProductResponse is the response for product data retrieval
// @Param ID int true "ID of the product"
// @Param Name string true "Name of the product"
// @Param Price int true "Price of the product"
// @Param Stock int true "Stock of the product"
// @Param Description string true "Description of the product"
// @Param Rating float64 true "Rating of the product"
// @Param Photos []string true "Photos of the product"
// @Param CategoryID int true "Category ID of the product"
// @Param Category Category true "Category of the product"
// @Param CreatedAt string true "Created at of the product"
// @Param UpdatedAt string true "Updated at of the product"
type ProductResponse struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Price       int               `json:"price"`
	Stock       int               `json:"stock"`
	Description string            `json:"description"`
	Rating      float64           `json:"rating"`
	Photos      []string          `json:"photos"`
	CategoryID  int               `json:"category_id"`
	Category    entities.Category `json:"category"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

func ProductFromEntities(product entities.Product, photos []entities.ProductPhoto) ProductResponse {

	photoUrls := make([]string, len(photos))
	for i, photo := range photos {
		photoUrls[i] = photo.UrlPhoto
	}

	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.Description,
		Rating:      product.Rating,
		Photos:      photoUrls,
		CategoryID:  product.CategoryID,
		Category:    product.Category,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
