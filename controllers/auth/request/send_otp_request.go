package request

// SendOTPRequest is the request for the send OTP endpoint
// @Description SendOTPRequest is the request for sending an OTP to the user's email
// @Param email body string true "Email address to send OTP to"
type SendOTPRequest struct {
	Email string `json:"email" form:"email" validate:"required,email"` 
}
