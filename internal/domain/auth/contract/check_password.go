package auth

type CheckPasswordRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CheckPasswordResponse struct {
}
