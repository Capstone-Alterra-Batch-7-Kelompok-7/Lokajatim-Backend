package request

type ChatbotRequest struct {
	Message string `json:"message" validate:"required"` // Ensure it's correctly marked as "message"
}
