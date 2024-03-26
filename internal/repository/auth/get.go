package auth

import (
	"context"
	"strings"

	authDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model"
	tokenRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *TokenRepository) Get(ctx context.Context, id string) (otp *authDomain.Token, err error) {
	var payload string
	payload, err = s.cache.Get(ctx, tokenRepository.GetTokenKey(id))
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrRepositoryAuth
		return
	}
	if (payload == "") || (strings.Compare(payload, "") == 0) {
		err = utils.ErrInvalidRefreshToken
		return
	}
	otp, err = tokenRepository.ToTokenDomain(payload)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrRepositoryAuth
		return
	}
	return
}
