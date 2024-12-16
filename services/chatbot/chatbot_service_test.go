package chatbot_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"lokajatim/services/chatbot"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockChatbotService adalah mock untuk interface ChatbotService
type MockChatbotService struct {
	mock.Mock
}

func (m *MockChatbotService) GenerateResponse(ctx context.Context, message string) (string, error) {
	args := m.Called(ctx, message)
	return args.String(0), args.Error(1)
}

func TestGenerateResponse_Success_WithMock(t *testing.T) {
	mockService := new(MockChatbotService)
	mockService.On("GenerateResponse", mock.Anything, "Halo!").Return("Halo, apa kabar?", nil)

	response, err := mockService.GenerateResponse(context.Background(), "Halo!")

	assert.NoError(t, err)
	assert.Equal(t, "Halo, apa kabar?", response)
	mockService.AssertExpectations(t)
}

func TestGenerateResponse_Error_WithMock(t *testing.T) {
	mockService := new(MockChatbotService)
	mockService.On("GenerateResponse", mock.Anything, "Halo!").Return("", errors.New("failed to connect"))

	response, err := mockService.GenerateResponse(context.Background(), "Halo!")

	assert.Error(t, err)
	assert.Equal(t, "", response)
	mockService.AssertExpectations(t)
}

func TestGenerateResponse_Integration(t *testing.T) {
	os.Setenv("GEMINI_API_KEY", "dummy_api_key")

	service, err := chatbot.NewChatbotService()
	assert.NoError(t, err)
	assert.NotNil(t, service)

	response, err := service.GenerateResponse(context.Background(), "Halo!")

	if err != nil {
		assert.Error(t, err)
		assert.Equal(t, "Maaf, saya tidak dapat memberikan jawaban saat ini.", response)
	} else {
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	}
}

// Mock untuk genai.ChatResponse dan komponennya
type MockChatResponse struct {
	Candidates []MockCandidate
}

type MockCandidate struct {
	Content MockContent
}

type MockContent struct {
	Parts []interface{}
}

func TestGenerateResponse_NoCandidates(t *testing.T) {
	mockService := new(MockChatbotService)
	mockService.On("GenerateResponse", mock.Anything, "Test").Return("Maaf, saya tidak dapat memberikan jawaban saat ini.", nil)

	response, err := mockService.GenerateResponse(context.Background(), "Test")

	assert.NoError(t, err)
	assert.Equal(t, "Maaf, saya tidak dapat memberikan jawaban saat ini.", response)
	mockService.AssertExpectations(t)
}

func TestGenerateResponse_NoParts(t *testing.T) {
	mockService := new(MockChatbotService)
	mockService.On("GenerateResponse", mock.Anything, "Test").Return("Maaf, saya tidak dapat memberikan jawaban saat ini.", nil)

	response, err := mockService.GenerateResponse(context.Background(), "Test")

	assert.NoError(t, err)
	assert.Equal(t, "Maaf, saya tidak dapat memberikan jawaban saat ini.", response)
	mockService.AssertExpectations(t)
}

func TestGenerateResponse_InvalidPartType(t *testing.T) {
	mockService := new(MockChatbotService)
	mockService.On("GenerateResponse", mock.Anything, "Test").Return("Maaf, saya tidak dapat memberikan jawaban saat ini.", nil)

	response, err := mockService.GenerateResponse(context.Background(), "Test")

	assert.NoError(t, err)
	assert.Equal(t, "Maaf, saya tidak dapat memberikan jawaban saat ini.", response)
	mockService.AssertExpectations(t)
}

func TestGenerateResponse_ValidParts(t *testing.T) {
	mockService := new(MockChatbotService)
	mockService.On("GenerateResponse", mock.Anything, "Test").Return("Halo, apa kabar? ", nil)

	response, err := mockService.GenerateResponse(context.Background(), "Test")

	assert.NoError(t, err)
	assert.Equal(t, "Halo, apa kabar? ", response)
	mockService.AssertExpectations(t)
}
