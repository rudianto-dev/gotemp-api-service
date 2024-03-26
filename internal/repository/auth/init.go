package auth

import (
	AuthInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth"
	"github.com/rudianto-dev/gotemp-sdk/pkg/cache"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type TokenRepository struct {
	logger *logger.Logger
	cache  *cache.DataSource
}

func NewAuthRepository(logger *logger.Logger, cache *cache.DataSource) (AuthInterface.Repository, error) {
	TokenRepo := &TokenRepository{
		logger: logger,
		cache:  cache,
	}
	return TokenRepo, nil
}
