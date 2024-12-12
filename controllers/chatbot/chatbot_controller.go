package chatbot

import (
	"log"
	"lokajatim/controllers/base"
	"lokajatim/controllers/chatbot/request"
	services "lokajatim/services/chatbot"
	"strings"

	"github.com/labstack/echo/v4"
)

type ChatbotController struct {
	AIService *services.AIService
}

func NewChatbotController(aiService *services.AIService) *ChatbotController {
	return &ChatbotController{
		AIService: aiService,
	}
}

func (c *ChatbotController) ChatbotController(ctx echo.Context) error {
	var requestBody request.RequestBody

	// Bind the incoming JSON body to the struct
	if err := ctx.Bind(&requestBody); err != nil {
		// Log the error for debugging purposes
		log.Printf("Failed to bind request body: %v", err)
		return base.ErrorResponse(ctx, err, map[string]string{"error": "Invalid request body"})
	}

	// Ensure the message is not empty and sanitize input
	requestBody.Message = strings.TrimSpace(requestBody.Message)
	if requestBody.Message == "" {
		return base.ErrorResponse(ctx, nil, map[string]string{"error": "message parameter is required"})
	}

	// Generate response from the service
	response, err := c.AIService.GenerateResponse(ctx.Request().Context(), requestBody.Message)
	if err != nil {
		// Log the error for debugging purposes
		log.Printf("Failed to generate response: %v", err)
		return base.ErrorResponse(ctx, err, map[string]string{"error": "Failed to generate response"})
	}

	// Return the generated response in a structured format
	return base.SuccesResponse(ctx, map[string]string{
		"response": response,
	})
}
