package auth

import (
	"github.com/go-redis/redis"
	AuthInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type TokenRepository struct {
	logger *logger.Logger
	redis  *redis.Client
}

func NewAuthRepository(logger *logger.Logger, redis *redis.Client) (AuthInterface.Repository, error) {
	TokenRepo := &TokenRepository{
		logger: logger,
		redis:  redis,
	}
	return TokenRepo, nil
}
