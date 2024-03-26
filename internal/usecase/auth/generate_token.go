package auth

import (
	"context"

	authDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model"
	authType "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model/type"
	authRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/repository"
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
	"github.com/rudianto-dev/gotemp-sdk/pkg/token"
)

func (s *AuthUseCase) GenerateToken(ctx context.Context, user *userDomain.User) (newToken *authDomain.Token, err error) {
	newToken = authDomain.NewToken(authType.Token{UserID: user.ID})
	newToken.Value, newToken.ExpiredAt, err = s.jwt.Create(token.Payload{ID: newToken.ID, UserID: newToken.UserID, RoleType: int8(user.RoleType)})
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
	return
}
