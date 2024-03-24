package module

import (
	"github.com/go-redis/redis"
	"github.com/rudianto-dev/gotemp-api-service/cmd/configuration"
	otpInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp"
	userInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	optRepository "github.com/rudianto-dev/gotemp-api-service/internal/repository/otp"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/repository/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/database"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type Module struct {
	Infra          *Service
	UserRepository userInterface.Repository
	OTPRepository  otpInterface.Repository
}

type Service struct {
	Config *configuration.ConfigurationSchema
	Logger *logger.Logger
	DB     *database.DB
	Redis  *redis.Client
}

func NewModule(infra *Service) *Module {
	userRepository, err := userRepository.NewUserRepository(infra.Logger, infra.DB)
	if err != nil {
		infra.Logger.Panicf("error init user repository, %v", err)
	}
	otpRepository, err := optRepository.NewOTPRepository(infra.Logger, infra.Redis)
	if err != nil {
		infra.Logger.Panicf("error init user repository, %v", err)
	}

	module := &Module{
		Infra:          infra,
		UserRepository: userRepository,
		OTPRepository:  otpRepository,
	}
	return module
}
