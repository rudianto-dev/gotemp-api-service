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
}

type HandlerAPI interface {
	CheckAccount(w http.ResponseWriter, r *http.Request)
}
