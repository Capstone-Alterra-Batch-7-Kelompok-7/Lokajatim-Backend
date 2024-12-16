package like_test

import (
	"errors"
	"lokajatim/entities"
	"lokajatim/services/like"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockLikeRepository untuk LikeRepoInterface
type MockLikeRepository struct {
	mock.Mock
}

func (m *MockLikeRepository) CreateLike(like entities.Like) (entities.Like, error) {
	args := m.Called(like)
	return args.Get(0).(entities.Like), args.Error(1)
}

func (m *MockLikeRepository) DeleteLike(articleID, userID int) error {
	args := m.Called(articleID, userID)
	return args.Error(0)
}

func (m *MockLikeRepository) GetLikesByArticleID(articleID int) ([]entities.Like, error) {
	args := m.Called(articleID)
	return args.Get(0).([]entities.Like), args.Error(1)
}

func (m *MockLikeRepository) CountLikesByArticleID(articleID int) (int, error) {
	args := m.Called(articleID)
	return args.Int(0), args.Error(1)
}

func (m *MockLikeRepository) IsUserLiked(articleID, userID int) (bool, error) {
	args := m.Called(articleID, userID)
	return args.Bool(0), args.Error(1)
}

// Test LikeArticle sukses
func TestLikeArticle_Success(t *testing.T) {
	mockRepo := new(MockLikeRepository)
	likeService := like.NewLikeService(mockRepo)

	articleID := 1
	userID := 1

	mockRepo.On("IsUserLiked", articleID, userID).Return(false, nil)
	mockRepo.On("CreateLike", mock.Anything).Return(entities.Like{ArticleID: articleID, UserID: userID}, nil)

	result, err := likeService.LikeArticle(articleID, userID)

	assert.NoError(t, err)
	assert.Equal(t, articleID, result.ArticleID)
	assert.Equal(t, userID, result.UserID)
	mockRepo.AssertExpectations(t)
}

// Test LikeArticle gagal (sudah di-like)
func TestLikeArticle_AlreadyLiked(t *testing.T) {
	mockRepo := new(MockLikeRepository)
	likeService := like.NewLikeService(mockRepo)

	articleID := 1
	userID := 1

	mockRepo.On("IsUserLiked", articleID, userID).Return(true, nil)

	result, err := likeService.LikeArticle(articleID, userID)

	assert.Error(t, err)
	assert.EqualError(t, err, "user already liked this article")
	assert.Equal(t, entities.Like{}, result)
	mockRepo.AssertExpectations(t)
}

// Test UnlikeArticle sukses
func TestUnlikeArticle_Success(t *testing.T) {
	mockRepo := new(MockLikeRepository)
	likeService := like.NewLikeService(mockRepo)

	articleID := 1
	userID := 1

	mockRepo.On("IsUserLiked", articleID, userID).Return(true, nil)
	mockRepo.On("DeleteLike", articleID, userID).Return(nil)

	err := likeService.UnlikeArticle(articleID, userID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// Test UnlikeArticle gagal (belum di-like)
func TestUnlikeArticle_NotLiked(t *testing.T) {
	mockRepo := new(MockLikeRepository)
	likeService := like.NewLikeService(mockRepo)

	articleID := 1
	userID := 1

	mockRepo.On("IsUserLiked", articleID, userID).Return(false, nil)

	err := likeService.UnlikeArticle(articleID, userID)

	assert.Error(t, err)
	assert.EqualError(t, err, "user has not liked this article")
	mockRepo.AssertExpectations(t)
}

// Test GetLikesByArticle sukses
func TestGetLikesByArticle_Success(t *testing.T) {
	mockRepo := new(MockLikeRepository)
	likeService := like.NewLikeService(mockRepo)

	articleID := 1
	expectedLikes := []entities.Like{
		{ArticleID: articleID, UserID: 1},
		{ArticleID: articleID, UserID: 2},
	}

	mockRepo.On("GetLikesByArticleID", articleID).Return(expectedLikes, nil)

	result, err := likeService.GetLikesByArticle(articleID)

	assert.NoError(t, err)
	assert.Equal(t, expectedLikes, result)
	mockRepo.AssertExpectations(t)
}

// Test CountLikes sukses
func TestCountLikes_Success(t *testing.T) {
	mockRepo := new(MockLikeRepository)
	likeService := like.NewLikeService(mockRepo)

	articleID := 1
	expectedCount := 5

	mockRepo.On("CountLikesByArticleID", articleID).Return(expectedCount, nil)

	result, err := likeService.CountLikes(articleID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCount, result)
	mockRepo.AssertExpectations(t)
}

// Test IsUserLikedArticle sukses
func TestIsUserLikedArticle_Success(t *testing.T) {
	mockRepo := new(MockLikeRepository)
	likeService := like.NewLikeService(mockRepo)

	articleID := 1
	userID := 1

	mockRepo.On("IsUserLiked", articleID, userID).Return(true, nil)

	result, err := likeService.IsUserLikedArticle(articleID, userID)

	assert.NoError(t, err)
	assert.True(t, result)
	mockRepo.AssertExpectations(t)
}

// Test IsUserLikedArticle gagal
func TestIsUserLikedArticle_Error(t *testing.T) {
	mockRepo := new(MockLikeRepository)
	likeService := like.NewLikeService(mockRepo)

	articleID := 1
	userID := 1

	mockRepo.On("IsUserLiked", articleID, userID).Return(false, errors.New("failed to check like status"))

	result, err := likeService.IsUserLikedArticle(articleID, userID)

	assert.Error(t, err)
	assert.False(t, result)
	assert.EqualError(t, err, "failed to check like status")
	mockRepo.AssertExpectations(t)
}
