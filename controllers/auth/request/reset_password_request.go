package request

// ResetPasswordRequest is the request for the reset password endpoint
// @Description ResetPasswordRequest is the request for the reset password endpoint
// @Param email body string true "Email of the user"
// @Param otp body string true "OTP sent to the user for password reset"
// @Param new_password body string true "New password to be set for the user"
type ResetPasswordRequest struct {
	Email       string `json:"email"`      
	OTP         string `json:"otp"`         
	NewPassword string `json:"new_password"` 
}
