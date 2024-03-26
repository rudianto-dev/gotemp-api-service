package auth

import (
	"context"

	authContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/contract"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *AuthUseCase) RefreshToken(ctx context.Context, req authContract.RefreshTokenRequest) (res *authContract.RefreshTokenResponse, err error) {
	oldToken, err := s.authRepository.Get(ctx, req.RefreshTokenID)
	if err != nil {
		return
	}
	user, err := s.userRepository.GetByID(ctx, oldToken.UserID)
	if err != nil {
		if err == utils.ErrNotFound {
			err = utils.ErrInvalidRefreshToken
		}
		return
	}

	newToken, err := s.GenerateToken(ctx, user)
	if err != nil {
		return
	}
	_ = s.authRepository.Delete(ctx, req.RefreshTokenID)
	res = &authContract.RefreshTokenResponse{
		Token:          newToken.Value,
		RefreshTokenID: newToken.ID,
		ExpiredAt:      newToken.ExpiredAt,
	}
	return
}
