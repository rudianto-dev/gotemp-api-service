package client

import (
	"context"
	"strings"

	clientRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *ClientRepository) GetCache(ctx context.Context, clientID string) (clientSecret string, err error) {
	var payload string
	payload, err = s.cache.Get(ctx, clientRepository.GetClientKey(clientID))
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrRepositoryClient
		return
	}
	if (payload == "") || (strings.Compare(payload, "") == 0) {
		err = utils.ErrNotFound
		return
	}
	payload = clientSecret
	return
}
