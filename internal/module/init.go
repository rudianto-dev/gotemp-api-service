package module

import (
	"github.com/go-redis/redis"
	"github.com/rudianto-dev/gotemp-api-service/cmd/configuration"
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	userRepo "github.com/rudianto-dev/gotemp-api-service/internal/repository/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/database"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type Module struct {
	infra    *Service
	userRepo userDomain.IRepository
}

type Service struct {
	Config *configuration.ConfigurationSchema
	Logger *logger.Logger
	DB     *database.DB
	Redis  *redis.Client
}

func NewModule(infra *Service) *Module {
	userRepo, err := userRepo.NewUserRepository(infra.Logger, infra.DB)
	if err != nil {
		infra.Logger.Panicf("error init user repository, %v", err)
	}

	module := &Module{
		infra:    infra,
		userRepo: userRepo,
	}
	return module
}
