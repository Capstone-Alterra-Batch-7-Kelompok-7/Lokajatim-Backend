package request

type RequestBody struct {
	Message string `json:"message" validate:"required"` // Ensure it's correctly marked as "message"
}
