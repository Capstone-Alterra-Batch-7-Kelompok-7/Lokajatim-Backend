package product_test

import (
	"lokajatim/entities"
	"lokajatim/services/product"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetProducts() ([]entities.Product, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return []entities.Product{}, args.Error(1)
	}
	return args.Get(0).([]entities.Product), args.Error(1)
}

func (m *MockProductRepository) GetProductByID(id int) (entities.Product, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Product), args.Error(1)
}

func (m *MockProductRepository) GetBestProductsPrice() ([]entities.Product, error) {
	args := m.Called()
	return args.Get(0).([]entities.Product), args.Error(1)
}

func (m *MockProductRepository) CreateProduct(product entities.Product) (entities.Product, error) {
	args := m.Called(product)
	return args.Get(0).(entities.Product), args.Error(1)
}

func (m *MockProductRepository) UpdateProduct(id int, product entities.Product) (entities.Product, error) {
	args := m.Called(id, product)
	return args.Get(0).(entities.Product), args.Error(1)
}

func (m *MockProductRepository) DeleteProduct(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockProductRepository) CreateProductPhotos(photos []entities.ProductPhoto) error {
	args := m.Called(photos)
	return args.Error(0)
}

func (m *MockProductRepository) GetProductPhotos(productID int) ([]entities.ProductPhoto, error) {
	args := m.Called(productID)
	return args.Get(0).([]entities.ProductPhoto), args.Error(1)
}

func (m *MockProductRepository) UpdateProductPhotos(productID int, photos []entities.ProductPhoto) error {
	args := m.Called(productID, photos)
	return args.Error(0)
}

func (m *MockProductRepository) DeleteProductPhotos(productID int) error {
	args := m.Called(productID)
	return args.Error(0)
}

func (m *MockProductRepository) BulkInsert(products []entities.Product) error {
	args := m.Called(products)
	return args.Error(0)
}

func (m *MockProductRepository) BulkInsertPhotos(photos []entities.ProductPhoto) error {
	args := m.Called(photos)
	return args.Error(0)
}

func TestGetProducts(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := product.NewProductService(mockRepo)

	mockProducts := []entities.Product{
		{ID: 1, Name: "Product 1", Price: 100},
		{ID: 2, Name: "Product 2", Price: 200},
	}
	mockRepo.On("GetProducts").Return(mockProducts, nil)

	products, err := service.GetProducts()

	assert.NoError(t, err)
	assert.Equal(t, mockProducts, products)

	mockRepo.AssertExpectations(t)
}

func TestGetProductsError(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := product.NewProductService(mockRepo)

	// Return an error from the mock
	mockRepo.On("GetProducts").Return(nil, assert.AnError)

	// Call the method
	products, err := service.GetProducts()

	// Check that the error is returned
	assert.Error(t, err)
	assert.Empty(t, products) // Expecting nil as products when an error occurs

	mockRepo.AssertExpectations(t)
}

func TestGetProductsEmptyList(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := product.NewProductService(mockRepo)

	mockRepo.On("GetProducts").Return([]entities.Product{}, nil)  // Empty list

	products, err := service.GetProducts()

	assert.NoError(t, err)
	assert.Empty(t, products)

	mockRepo.AssertExpectations(t)
}

func TestGetProductByID(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := product.NewProductService(mockRepo)

	mockProduct := entities.Product{ID: 1, Name: "Product 1", Price: 100}
	mockRepo.On("GetProductByID", 1).Return(mockProduct, nil)

	product, err := service.GetProductByID(1)

	assert.NoError(t, err)
	assert.Equal(t, mockProduct, product)

	mockRepo.AssertExpectations(t)
}

func TestCreateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := product.NewProductService(mockRepo)

	mockProduct := entities.Product{Name: "Product 1", Price: 100}
	createdProduct := entities.Product{ID: 1, Name: "Product 1", Price: 100}
	mockRepo.On("CreateProduct", mockProduct).Return(createdProduct, nil)

	product, err := service.CreateProduct(mockProduct)

	assert.NoError(t, err)
	assert.Equal(t, createdProduct, product)

	mockRepo.AssertExpectations(t)
}

func TestCreateProductPhotos(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := product.NewProductService(mockRepo)

	photos := []entities.ProductPhoto{
		{ProductID: 1, UrlPhoto: "url1"},
		{ProductID: 1, UrlPhoto: "url2"},
	}

	mockRepo.On("CreateProductPhotos", photos).Return(nil)  // Simulate successful photo creation

	err := service.CreateProductPhotos(photos)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := product.NewProductService(mockRepo)

	mockProduct := entities.Product{Name: "Updated Product", Price: 150}
	updatedProduct := entities.Product{ID: 1, Name: "Updated Product", Price: 150}
	mockRepo.On("UpdateProduct", 1, mockProduct).Return(updatedProduct, nil)

	product, err := service.UpdateProduct(1, mockProduct)

	assert.NoError(t, err)
	assert.Equal(t, updatedProduct, product)

	mockRepo.AssertExpectations(t)
}

func TestUpdateProductPhotos(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := product.NewProductService(mockRepo)

	photos := []entities.ProductPhoto{
		{ProductID: 1, UrlPhoto: "new_url1"},
	}

	mockRepo.On("UpdateProductPhotos", 1, photos).Return(nil)  // Simulate successful photo update

	err := service.UpdateProductPhotos(1, photos)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := product.NewProductService(mockRepo)

	mockRepo.On("DeleteProduct", 1).Return(nil)

	err := service.DeleteProduct(1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetBestProductsPrice(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := product.NewProductService(mockRepo)

	mockProducts := []entities.Product{
		{ID: 1, Name: "Best Product", Price: 50000},
		{ID: 2, Name: "Another Product", Price: 40000},
	}
	mockRepo.On("GetBestProductsPrice").Return(mockProducts, nil)

	products, err := service.GetBestProductsPrice()

	assert.NoError(t, err)
	assert.Equal(t, mockProducts, products)

	mockRepo.AssertExpectations(t)
}

func TestGetBestProductsPriceError(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := product.NewProductService(mockRepo)

	mockRepo.On("GetBestProductsPrice").Return([]entities.Product{}, assert.AnError)

	products, err := service.GetBestProductsPrice()

	assert.Error(t, err)
	assert.Empty(t, products)

	mockRepo.AssertExpectations(t)
}

