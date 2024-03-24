package auth

import (
	"context"
	"net/http"

	authContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/contract"
)

type Repository interface {
}

type UseCase interface {
	CheckAccount(ctx context.Context, req authContract.CheckAccountRequest) (*authContract.CheckAccountResponse, error)
	ResetPassword(ctx context.Context, req authContract.ResetPasswordRequest) (*authContract.ResetPasswordResponse, error)
}

type HandlerAPI interface {
	CheckAccount(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
}
