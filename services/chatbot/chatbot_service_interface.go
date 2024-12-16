package chatbot

import (
	"context"
)

type ChatbotService interface {
	GenerateResponse(ctx context.Context, message string) (string, error)
}
