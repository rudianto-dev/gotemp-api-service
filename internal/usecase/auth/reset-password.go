package auth

import (
	"context"

	authContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/contract"
)

func (s *AuthUseCase) ResetPassword(ctx context.Context, req authContract.ResetPasswordRequest) (res *authContract.ResetPasswordResponse, err error) {
	// user, err := s.userRepo.GetByUsername(ctx, req.Username)
	// if err != nil {
	// 	return
	// }
	res = &authContract.ResetPasswordResponse{}
	return
}
