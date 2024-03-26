package auth

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token          string `json:"token" validate:"required"`
	RefreshTokenID string `json:"refresh_token_id" validate:"required"`
	ExpiredAt      int64  `json:"expired_at" validate:"required"`
}
