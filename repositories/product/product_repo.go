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
	result := r.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *ProductRepositoryImpl) GetProductByID(id int) (entities.Product, error) {
	var product entities.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return entities.Product{}, result.Error
	}
	return product, nil
}

func (r *ProductRepositoryImpl) CreateProduct(product entities.Product) (entities.Product, error) {
	result := r.db.Create(&product)
	if result.Error != nil {
		return entities.Product{}, result.Error
	}
	return product, nil
}

func (r *ProductRepositoryImpl) UpdateProduct(product entities.Product) (entities.Product, error) {
	result := r.db.Save(&product)
	if result.Error != nil {
		return entities.Product{}, result.Error
	}
	return product, nil
}

func (r *ProductRepositoryImpl) DeleteProduct(id int) error {
	var product entities.Product
	if err := r.db.Delete(&product, id).Error; err != nil {
		return err
	}
	return nil
}