package auth

import (
	"context"

	authContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/contract"
	authDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model"
	authType "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *AuthUseCase) Login(ctx context.Context, req authContract.LoginRequest) (res *authContract.LoginResponse, err error) {
	user, err := s.userRepository.GetByUsername(ctx, req.Username)
	if err != nil {
		if err == utils.ErrNotFound {
			err = utils.ErrInvalidCredential
		}
		return
	}

	newAuth := authDomain.New(authType.Credential{
		Username: user.Username,
		Password: user.Password,
	})
	err = newAuth.CheckPassword(req.Password)
	if err != nil {
		err = utils.ErrInvalidCredential
		return
	}

	newToken, err := s.GenerateToken(ctx, user)
	if err != nil {
		return
	}
	res = &authContract.LoginResponse{
		Token:          newToken.Value,
		RefreshTokenID: newToken.ID,
		ExpiredAt:      newToken.ExpiredAt,
	}
	return
}
