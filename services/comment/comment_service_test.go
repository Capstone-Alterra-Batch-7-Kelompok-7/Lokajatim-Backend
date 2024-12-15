package comment_test

import (
	"errors"
	"lokajatim/entities"
	"lokajatim/services/comment"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCommentRepository untuk CommentRepositoryInterface
type MockCommentRepository struct {
	mock.Mock
}

func (m *MockCommentRepository) GetCommentByID(id int) (entities.Comment, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Comment), args.Error(1)
}

func (m *MockCommentRepository) GetCommentsByArticleID(articleID int) ([]entities.Comment, error) {
	args := m.Called(articleID)
	return args.Get(0).([]entities.Comment), args.Error(1)
}

func (m *MockCommentRepository) CreateComment(comment entities.Comment) (entities.Comment, error) {
	args := m.Called(comment)
	return args.Get(0).(entities.Comment), args.Error(1)
}

func (m *MockCommentRepository) DeleteComment(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Test GetCommentByID sukses
func TestGetCommentByID_Success(t *testing.T) {
	mockRepo := new(MockCommentRepository)
	commentService := comment.NewCommentService(mockRepo)

	expectedComment := entities.Comment{ID: 1, ArticleID: 1, Comment: "This is a test comment"}

	mockRepo.On("GetCommentByID", 1).Return(expectedComment, nil)

	result, err := commentService.GetCommentByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedComment, result)
	mockRepo.AssertExpectations(t)
}

// Test GetCommentByID error
func TestGetCommentByID_Error(t *testing.T) {
	mockRepo := new(MockCommentRepository)
	commentService := comment.NewCommentService(mockRepo)

	mockRepo.On("GetCommentByID", 1).Return(entities.Comment{}, errors.New("comment not found"))

	result, err := commentService.GetCommentByID(1)

	assert.Error(t, err)
	assert.Equal(t, entities.Comment{}, result)
	mockRepo.AssertExpectations(t)
}

// Test GetCommentsByArticleID sukses
func TestGetCommentsByArticleID_Success(t *testing.T) {
	mockRepo := new(MockCommentRepository)
	commentService := comment.NewCommentService(mockRepo)

	expectedComments := []entities.Comment{
		{ID: 1, ArticleID: 1, Comment: "Comment 1"},
		{ID: 2, ArticleID: 1, Comment: "Comment 2"},
	}

	mockRepo.On("GetCommentsByArticleID", 1).Return(expectedComments, nil)

	result, err := commentService.GetCommentsByArticleID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedComments, result)
	mockRepo.AssertExpectations(t)
}

// Test GetCommentsByArticleID error
func TestGetCommentsByArticleID_Error(t *testing.T) {
	mockRepo := new(MockCommentRepository)
	commentService := comment.NewCommentService(mockRepo)

	mockRepo.On("GetCommentsByArticleID", 1).Return(nil, errors.New("failed to fetch comments"))

	result, err := commentService.GetCommentsByArticleID(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// Test CreateComment sukses
func TestCreateComment_Success(t *testing.T) {
	mockRepo := new(MockCommentRepository)
	commentService := comment.NewCommentService(mockRepo)

	newComment := entities.Comment{ArticleID: 1, Comment: "New comment"}
	createdComment := entities.Comment{ID: 1, ArticleID: 1, Comment: "New comment"}

	mockRepo.On("CreateComment", newComment).Return(createdComment, nil)

	result, err := commentService.CreateComment(newComment)

	assert.NoError(t, err)
	assert.Equal(t, createdComment, result)
	mockRepo.AssertExpectations(t)
}

// Test CreateComment error
func TestCreateComment_Error(t *testing.T) {
	mockRepo := new(MockCommentRepository)
	commentService := comment.NewCommentService(mockRepo)

	newComment := entities.Comment{ArticleID: 1, Comment: "New comment"}

	mockRepo.On("CreateComment", newComment).Return(entities.Comment{}, errors.New("failed to create comment"))

	result, err := commentService.CreateComment(newComment)

	assert.Error(t, err)
	assert.Equal(t, entities.Comment{}, result)
	mockRepo.AssertExpectations(t)
}

// Test DeleteComment sukses
func TestDeleteComment_Success(t *testing.T) {
	mockRepo := new(MockCommentRepository)
	commentService := comment.NewCommentService(mockRepo)

	mockRepo.On("DeleteComment", 1).Return(nil)

	err := commentService.DeleteComment(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// Test DeleteComment error
func TestDeleteComment_Error(t *testing.T) {
	mockRepo := new(MockCommentRepository)
	commentService := comment.NewCommentService(mockRepo)

	mockRepo.On("DeleteComment", 1).Return(errors.New("failed to delete comment"))

	err := commentService.DeleteComment(1)

	assert.Error(t, err)
	assert.EqualError(t, err, "failed to delete comment")
	mockRepo.AssertExpectations(t)
}
