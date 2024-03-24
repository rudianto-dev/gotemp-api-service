package auth

import (
	"context"

	authContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/contract"
	authDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model"
	authType "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model/type"
	userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/token"
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
	token, expiredAt, err := s.jwt.Create(token.Payload{ID: user.ID, RoleType: int8(user.RoleType)})
	if err != nil {
		return
	}
	_ = s.otpRepository.DeleteVerification(ctx, req.OtpVerificationID)
	res = &authContract.ResetPasswordResponse{
		Token:     token,
		ExpiredAt: expiredAt}
	return
}
