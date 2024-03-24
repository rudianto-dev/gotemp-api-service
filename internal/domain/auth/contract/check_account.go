package auth

import (
	userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"
)

type CheckAccountRequest struct {
	Username string `json:"username" validate:"required"`
}

type CheckAccountResponse struct {
	Username string          `json:"username" validate:"required"`
	Status   userType.Status `json:"status" validate:"required"`
}
