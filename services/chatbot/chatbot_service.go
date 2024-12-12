package services

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

func (s *AIService) GenerateResponse(ctx context.Context, message string) (string, error) {
	// Tambahkan pesan pengguna ke sesi chat secara manual
	s.cs.History = append(s.cs.History, genai.NewUserContent(genai.Text(message)))

	// Kirim pesan ke model
	resp, err := s.cs.SendMessage(ctx, genai.Text(message))
	if err != nil {
		return "", fmt.Errorf("failed to generate response: %w", err)
	}

	// Periksa respons dan ekstrak teks
	if len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil {
		response := ""
		for _, part := range resp.Candidates[0].Content.Parts {
			if text, ok := part.(genai.Text); ok {
				response += string(text) + " "
			}
		}
		return response, nil
	}

	return "Maaf, saya tidak dapat memberikan jawaban saat ini.", nil
}

