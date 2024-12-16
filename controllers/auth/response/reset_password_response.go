package response

// ResetPasswordResponse is the response for the reset password endpoint
// @Description ResetPasswordResponse contains the message confirming the result of the password reset action
// @Success 200 {object} response.ResetPasswordResponse "Password reset successfully"
// @Failure 400 {object} base.BaseResponse "Invalid OTP or email"
// @Failure 500 {object} base.BaseResponse "Internal server error"
type ResetPasswordResponse struct {
	Message string `json:"message"`
}
