package chatbot

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIService struct {
	client *genai.Client
	model  *genai.GenerativeModel
	cs     *genai.ChatSession
}

func NewChatbotService() (*AIService, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return nil, fmt.Errorf("failed to create Generative AI client: %w", err)
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	cs := model.StartChat()

	return &AIService{
		client: client,
		model:  model,
		cs:     cs,
	}, nil
}

func (c *AIService) GenerateResponse(ctx context.Context, message string) (string, error) {
	c.cs.History = append(c.cs.History, genai.NewUserContent(genai.Text(message)))

	resp, err := c.cs.SendMessage(ctx, genai.Text(message))
	if err != nil || resp == nil || len(resp.Candidates) == 0 {
		return "Maaf, saya tidak dapat memberikan jawaban saat ini.", nil
	}

	response := ""
	if len(resp.Candidates[0].Content.Parts) > 0 {
		for _, part := range resp.Candidates[0].Content.Parts {
			if text, ok := part.(genai.Text); ok {
				response += string(text) + " "
			}
		}
	}

	if response == "" {
		return "Maaf, saya tidak dapat memberikan jawaban saat ini.", nil
	}

	return response, nil
}
