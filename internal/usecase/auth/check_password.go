package auth

import (
	"context"

	authContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/contract"
	authDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model"
	authType "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *AuthUseCase) CheckPassword(ctx context.Context, req authContract.CheckPasswordRequest) (res *authContract.CheckPasswordResponse, err error) {
	user, err := s.userRepository.GetByID(ctx, req.UserID)
	if err != nil {
		return
	}
	newAuth := authDomain.New(authType.Credential{
		Username: user.Username,
		Password: user.Password,
	})
	err = newAuth.CheckPassword(req.Password)
	if err != nil {
		err = utils.ErrInvalidPassword
		return
	}
	res = &authContract.CheckPasswordResponse{}
	return
}
