package auth

import (
	"context"

	tokenRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *TokenRepository) Delete(ctx context.Context, ID string) (err error) {
	_, err = s.redis.Del(tokenRepository.GetTokenKey(ID)).Result()
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrRepositoryOTP
		return
	}
	return
}
