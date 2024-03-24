package auth

import (
	"context"

	authContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/contract"
)

func (s *AuthUseCase) CheckAccount(ctx context.Context, req authContract.CheckAccountRequest) (res *authContract.CheckAccountResponse, err error) {
	user, err := s.userRepository.GetByUsername(ctx, req.Username)
	if err != nil {
		return
	}
	res = &authContract.CheckAccountResponse{
		Username: req.Username,
		Status:   user.Status,
	}
	return
}
