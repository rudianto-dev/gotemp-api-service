package client

import (
	"context"

	clientRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *ClientRepository) SaveCache(ctx context.Context, req *clientRepository.CacheEntity) (err error) {
	err = s.cache.Save(ctx, req.Key, req.Payload, req.TTL)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrRepositoryClient
		return
	}
	return
}
