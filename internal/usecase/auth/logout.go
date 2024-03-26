package auth

import (
	"context"

	authContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/contract"
)

func (s *AuthUseCase) Logout(ctx context.Context, req authContract.LogoutRequest) (res *authContract.LogoutResponse, err error) {
	err = s.authRepository.Delete(ctx, req.TokenID)
	if err != nil {
		return
	}
	res = &authContract.LogoutResponse{}
	return
}
