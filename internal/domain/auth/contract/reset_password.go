package auth

type ResetPasswordRequest struct {
	Password          string `json:"password" validate:"required"`
	OtpVerificationID string `json:"otp_verification_id" validate:"required"`
}

type ResetPasswordResponse struct {
}
