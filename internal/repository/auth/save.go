package auth

import (
	"context"

	authRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *TokenRepository) Save(ctx context.Context, req *authRepository.TokenEntity) (err error) {
	_, err = s.redis.Set(req.Key, req.Payload, req.TTL).Result()
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrRepositoryOTP
		return
	}
	return
}
