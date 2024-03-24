package auth

import (
	"context"
	"net/http"

	authContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/contract"
	authRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/repository"
)

type Repository interface {
	Save(ctx context.Context, req *authRepository.TokenEntity) error
	Delete(ctx context.Context, id string) error
}

type UseCase interface {
	CheckAccount(ctx context.Context, req authContract.CheckAccountRequest) (*authContract.CheckAccountResponse, error)
	CheckPassword(ctx context.Context, req authContract.CheckPasswordRequest) (*authContract.CheckPasswordResponse, error)
	ResetPassword(ctx context.Context, req authContract.ResetPasswordRequest) (*authContract.ResetPasswordResponse, error)

	Login(ctx context.Context, req authContract.LoginRequest) (*authContract.LoginResponse, error)
}

type HandlerAPI interface {
	CheckAccount(w http.ResponseWriter, r *http.Request)
	CheckPassword(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)

	Login(w http.ResponseWriter, r *http.Request)
}
