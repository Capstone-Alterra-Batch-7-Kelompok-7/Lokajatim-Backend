package response

// SendOTPResponse is the response for the send OTP endpoint
// @Description SendOTPResponse contains the message confirming the result of sending an OTP
// @Success 200 {object} response.SendOTPResponse "OTP sent successfully"
// @Failure 400 {object} base.BaseResponse "Invalid email address"
// @Failure 500 {object} base.BaseResponse "Internal server error"
type SendOTPResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
