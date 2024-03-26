package auth

import (
	"context"

	authContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/contract"
	authDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model"
	authType "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model/type"
	userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"
)

func (s *AuthUseCase) ResetPassword(ctx context.Context, req authContract.ResetPasswordRequest) (res *authContract.ResetPasswordResponse, err error) {
	verification, err := s.otpRepository.GetVerification(ctx, req.OtpVerificationID)
	if err != nil {
		return
	}
	user, err := s.userRepository.GetByUsername(ctx, verification.Receiver)
	if err != nil {
		return
	}
	hash, err := authDomain.New(authType.Credential{
		Username: verification.Receiver,
		Password: req.Password,
	}).HashPassword()
	if err != nil {
		return
	}
	err = s.userRepository.UpdatePassword(ctx, nil, user.ID, hash)
	if err != nil {
		return
	}
	err = s.userRepository.UpdateStatus(ctx, nil, user.ID, userType.USER_STATUS_ACTIVE)
	if err != nil {
		return
	}

	newToken, err := s.GenerateToken(ctx, user)
	if err != nil {
		return
	}
	_ = s.otpRepository.DeleteVerification(ctx, req.OtpVerificationID)
	res = &authContract.ResetPasswordResponse{
		Token:          newToken.Value,
		RefreshTokenID: newToken.ID,
		ExpiredAt:      newToken.ExpiredAt,
	}
	return
}
