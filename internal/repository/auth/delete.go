package auth

import (
	"context"

	tokenRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *TokenRepository) Delete(ctx context.Context, ID string) (err error) {
	err = s.cache.Delete(ctx, tokenRepository.GetTokenKey(ID))
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrRepositoryAuth
		return
	}
	return
}
