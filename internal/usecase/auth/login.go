package auth

import (
	"context"

	authContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/contract"
	authDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model"
	authType "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model/type"
	authRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/token"
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

	newToken := authDomain.NewToken(authType.Token{})
	newToken.Value, newToken.ExpiredAt, err = s.jwt.Create(token.Payload{ID: newToken.ID, UserID: user.ID, RoleType: int8(user.RoleType)})
	if err != nil {
		return
	}
	tokenEntity, err := authRepository.ToTokenEntity(newToken)
	if err != nil {
		return
	}
	err = s.authRepository.Save(ctx, tokenEntity)
	if err != nil {
		return
	}
	res = &authContract.LoginResponse{
		Token:     newToken.Value,
		ExpiredAt: newToken.ExpiredAt,
	}
	return
}
