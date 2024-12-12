package chatbot

import (
	"log"
	services "lokajatim/services/chatbot"
	"net/http"
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

type RequestBody struct {
	Message string `json:"message" validate:"required"` // Ensure it's correctly marked as "message"
}

func (c *ChatbotController) ChatbotController(ctx echo.Context) error {
	var requestBody RequestBody

	// Bind the incoming JSON body to the struct
	if err := ctx.Bind(&requestBody); err != nil {
		// Log the error for debugging purposes
		log.Printf("Failed to bind request body: %v", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Ensure the message is not empty and sanitize input
	requestBody.Message = strings.TrimSpace(requestBody.Message)
	if requestBody.Message == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "message parameter is required"})
	}

	// Generate response from the service
	response, err := c.AIService.GenerateResponse(ctx.Request().Context(), requestBody.Message)
	if err != nil {
		// Log the error for debugging purposes
		log.Printf("Failed to generate response: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate response"})
	}

	// Return the generated response in a structured format
	return ctx.JSON(http.StatusOK, map[string]string{
		"response": response,
	})
}
