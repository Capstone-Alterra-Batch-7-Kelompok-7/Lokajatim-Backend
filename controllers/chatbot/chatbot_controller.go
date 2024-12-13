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

// @Summary Generate Chatbot Response
// @Tags Chatbot
// @Accept json
// @Produce json
// @Param request body request.ChatbotRequest true "Chatbot Request Body"
// @Success 200 {object} base.BaseResponse{data=map[string]string} "Chatbot Response"
// @Failure 400 {object} base.BaseResponse{data=map[string]string} "Invalid Request Body"
// @Failure 500 {object} base.BaseResponse{data=map[string]string} "Internal Server Error"
// @Router /chatbot [post]
func (c *ChatbotController) ChatbotController(ctx echo.Context) error {
	var requestBody request.ChatbotRequest

	if err := ctx.Bind(&requestBody); err != nil {
		log.Printf("Failed to bind request body: %v", err)
		return base.ErrorResponse(ctx, err, map[string]string{"error": "Invalid request body"})
	}

	requestBody.Message = strings.TrimSpace(requestBody.Message)
	if requestBody.Message == "" {
		return base.ErrorResponse(ctx, nil, map[string]string{"error": "message parameter is required"})
	}

	response, err := c.AIService.GenerateResponse(ctx.Request().Context(), requestBody.Message)
	if err != nil {
		log.Printf("Failed to generate response: %v", err)
		return base.ErrorResponse(ctx, err, map[string]string{"error": "Failed to generate response"})
	}

	return base.SuccesResponse(ctx, map[string]string{
		"response": response,
	})
}
