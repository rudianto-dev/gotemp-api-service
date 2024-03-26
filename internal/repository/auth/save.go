package auth

import (
	"context"

	authRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *TokenRepository) Save(ctx context.Context, req *authRepository.TokenEntity) (err error) {
	err = s.cache.Save(ctx, req.Key, req.Payload, req.TTL)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrRepositoryOTP
		return
	}
	return
}
