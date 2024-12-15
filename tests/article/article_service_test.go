package article_test

import (
	"errors"
	"lokajatim/entities"
	"lokajatim/services/article"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockArticleRepository struct {
	mock.Mock
}

func (m *MockArticleRepository) GetAll() ([]entities.Article, error) {
	args := m.Called()
	return args.Get(0).([]entities.Article), args.Error(1)
}

func (m *MockArticleRepository) GetByID(id int) (entities.Article, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Article), args.Error(1)
}

func (m *MockArticleRepository) Create(article entities.Article) (entities.Article, error) {
	args := m.Called(article)
	return args.Get(0).(entities.Article), args.Error(1)
}

func (m *MockArticleRepository) Update(id int, article entities.Article) (entities.Article, error) {
	args := m.Called(id, article)
	return args.Get(0).(entities.Article), args.Error(1)
}

func (m *MockArticleRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllArticles(t *testing.T) {
	mockRepo := new(MockArticleRepository)
	service := article.NewArticleService(mockRepo)

	articles := []entities.Article{
		{ID: 1, Title: "Article 1"},
		{ID: 2, Title: "Article 2"},
	}

	mockRepo.On("GetAll").Return(articles, nil)

	result, err := service.GetAllArticles()

	assert.NoError(t, err)
	assert.Equal(t, articles, result)
	mockRepo.AssertExpectations(t)
}

func TestGetArticleByID(t *testing.T) {
	mockRepo := new(MockArticleRepository)
	service := article.NewArticleService(mockRepo)

	article := entities.Article{ID: 1, Title: "Article 1"}

	mockRepo.On("GetByID", 1).Return(article, nil)

	result, err := service.GetArticleByID(1)

	assert.NoError(t, err)
	assert.Equal(t, article, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateArticle(t *testing.T) {
	mockRepo := new(MockArticleRepository)
	service := article.NewArticleService(mockRepo)

	newArticle := entities.Article{Title: "New Article"}
	createdArticle := entities.Article{ID: 1, Title: "New Article"}

	mockRepo.On("Create", newArticle).Return(createdArticle, nil)

	result, err := service.CreateArticle(newArticle)

	assert.NoError(t, err)
	assert.Equal(t, createdArticle, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateArticle(t *testing.T) {
	mockRepo := new(MockArticleRepository)
	service := article.NewArticleService(mockRepo)

	updatedArticle := entities.Article{Title: "Updated Article"}
	finalArticle := entities.Article{ID: 1, Title: "Updated Article"}

	mockRepo.On("Update", 1, updatedArticle).Return(finalArticle, nil)

	result, err := service.UpdateArticle(1, updatedArticle)

	assert.NoError(t, err)
	assert.Equal(t, finalArticle, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteArticle(t *testing.T) {
	mockRepo := new(MockArticleRepository)
	service := article.NewArticleService(mockRepo)

	mockRepo.On("Delete", 1).Return(nil)

	err := service.DeleteArticle(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetArticleByID_Error(t *testing.T) {
	mockRepo := new(MockArticleRepository)
	service := article.NewArticleService(mockRepo)

	mockRepo.On("GetByID", 1).Return(entities.Article{}, errors.New("article not found"))

	result, err := service.GetArticleByID(1)

	assert.Error(t, err)
	assert.Equal(t, entities.Article{}, result)
	mockRepo.AssertExpectations(t)
}
