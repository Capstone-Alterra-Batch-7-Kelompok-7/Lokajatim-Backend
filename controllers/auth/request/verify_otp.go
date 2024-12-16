package request

type VerifyOTPRequest struct {
    OTP string `json:"otp" validate:"required"`
}
