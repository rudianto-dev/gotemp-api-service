package auth

type RefreshTokenRequest struct {
	RefreshTokenID string `json:"refresh_token_id" validate:"required"`
}

type RefreshTokenResponse struct {
	Token          string `json:"token" validate:"required"`
	RefreshTokenID string `json:"refresh_token_id" validate:"required"`
	ExpiredAt      int64  `json:"expired_at" validate:"required"`
}
