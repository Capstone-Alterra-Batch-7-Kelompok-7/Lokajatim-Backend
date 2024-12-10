package request

import "lokajatim/entities"

// ProductRequest is the request for the Product endpoint
// @Description ProductRequest is the request for the Product endpoint
// @Param Name string true "Name of the product"
// @Param Price int true "Price of the product"
// @Param Stock int true "Stock of the product"
// @Param Description string true "Description of the product"
// @Param Photo string true "URL Photo of the product"
// @Param CategoryID int true "Category ID of the product"
type ProductRequest struct {
	Name        string   `json:"name" validate:"required"`
	Price       int      `json:"price" validate:"required"`
	Stock       int      `json:"stock" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Rating      float64  `json:"rating" validate:"required"`
	CategoryID  int      `json:"category_id" validate:"required"`
	Photos      []string `json:"photos" validate:"required"`
}

type PhotoRequest struct {
	UrlPhoto string `json:"url_photo" validate:"required"`
}

func (productRequest ProductRequest) ToEntities() (entities.Product, error) {
	product := entities.Product{
		Name:        productRequest.Name,
		Price:       productRequest.Price,
		Stock:       productRequest.Stock,
		Description: productRequest.Description,
		Rating:      productRequest.Rating,
		CategoryID:  productRequest.CategoryID,
	}
	return product, nil
}

func (productRequest ProductRequest) ToProductPhotos(productID int) ([]entities.ProductPhoto, error) {
	photos := make([]entities.ProductPhoto, len(productRequest.Photos))
	for i, url := range productRequest.Photos {
		photos[i] = entities.ProductPhoto{
			UrlPhoto:  url,
			ProductID: productID,
		}
	}
	return photos, nil
}
