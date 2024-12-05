package product

import (
	"lokajatim/entities"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryInterface {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) GetProducts() ([]entities.Product, error) {
	var products []entities.Product
	result := r.db.Preload("Category").Preload("Photos").Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *ProductRepositoryImpl) GetProductByID(id int) (entities.Product, error) {
	var product entities.Product
	result := r.db.Preload("Category").Preload("Photos").First(&product, id)
	if result.Error != nil {
		return entities.Product{}, result.Error
	}
	return product, nil
}

func (r *ProductRepositoryImpl) CreateProduct(product entities.Product) (entities.Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return entities.Product{}, err
	}

	var createdProduct entities.Product
	result := r.db.Preload("Category").Preload("Photos").First(&createdProduct, product.ID)
	if result.Error != nil {
		return entities.Product{}, result.Error
	}
	return createdProduct, nil
}

func (r *ProductRepositoryImpl) UpdateProduct(id int, product entities.Product) (entities.Product, error) {
	if err := r.db.Model(&entities.Product{}).Where("id = ?", id).Updates(product).Error; err != nil {
		return entities.Product{}, err
	}

	var updatedProduct entities.Product
	result := r.db.Preload("Category").Preload("Photos").First(&updatedProduct, id)
	if result.Error != nil {
		return entities.Product{}, result.Error
	}
	return updatedProduct, nil
}

func (r *ProductRepositoryImpl) DeleteProduct(id int) error {
	var product entities.Product
	if err := r.db.Preload("Photos").Delete(&product, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) CreateProductPhotos(photos []entities.ProductPhoto) error {
	result := r.db.Create(&photos)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ProductRepositoryImpl) GetProductPhotos(productID int) ([]entities.ProductPhoto, error) {
	var photos []entities.ProductPhoto
	result := r.db.Where("product_id = ?", productID).Find(&photos)
	if result.Error != nil {
		return nil, result.Error
	}
	return photos, nil
}

func (r *ProductRepositoryImpl) UpdateProductPhotos(productID int, photos []entities.ProductPhoto) error {
	if err := r.db.Where("product_id = ?", productID).Delete(&entities.ProductPhoto{}).Error; err != nil {
		return err
	}

	for i := range photos {
		photos[i].ProductID = productID
	}

	if err := r.db.Create(&photos).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) DeleteProductPhotos(productID int) error {
	if err := r.db.Where("product_id = ?", productID).Delete(&entities.ProductPhoto{}).Error; err != nil {
		return err
	}
	return nil
}
