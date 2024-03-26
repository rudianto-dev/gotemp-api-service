package client

import (
	"context"

	clientRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *ClientRepository) DeleteCache(ctx context.Context, clientID string) (err error) {
	err = s.cache.Delete(ctx, clientRepository.GetClientKey(clientID))
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrRepositoryClient
		return
	}
	return
}
