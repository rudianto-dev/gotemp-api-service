package auth

type ResetPasswordRequest struct {
	Password          string `json:"password" validate:"required"`
	OtpVerificationID string `json:"otp_verification_id" validate:"required"`
}

type ResetPasswordResponse struct {
	Token          string `json:"token" validate:"required"`
	RefreshTokenID string `json:"refresh_token_id" validate:"required"`
	ExpiredAt      int64  `json:"expired_at" validate:"required"`
}
