package auth

type LogoutRequest struct {
	TokenID string `json:"token_id" validate:"required"`
}

type LogoutResponse struct {
}
