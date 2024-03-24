package otp

type VerifyOTPRequest struct {
	Receiver string `json:"receiver" validate:"required"`
	Code     string `json:"code" validate:"required"`
}

type VerifyOTPResponse struct {
	OTPVerificationID string `json:"otp_verification_id"`
}
