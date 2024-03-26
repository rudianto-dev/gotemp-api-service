package module

import (
	"github.com/rudianto-dev/gotemp-api-service/cmd/configuration"
	authInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth"
	otpInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp"
	userInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	authRepository "github.com/rudianto-dev/gotemp-api-service/internal/repository/auth"
	otpRepository "github.com/rudianto-dev/gotemp-api-service/internal/repository/otp"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/repository/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/cache"
	"github.com/rudianto-dev/gotemp-sdk/pkg/database"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
	"github.com/rudianto-dev/gotemp-sdk/pkg/token"
)

type Module struct {
	Infra          *Service
	UserRepository userInterface.Repository
	OTPRepository  otpInterface.Repository
	AuthRepository authInterface.Repository
}

type Service struct {
	Config *configuration.ConfigurationSchema
	Logger *logger.Logger
	JWT    *token.JWT
	DB     *database.DB
	Cache  *cache.DataSource
}

func NewModule(infra *Service) *Module {
	userRepository, err := userRepository.NewUserRepository(infra.Logger, infra.DB)
	if err != nil {
		infra.Logger.Panicf("error init user repository, %v", err)
	}
	authRepository, err := authRepository.NewAuthRepository(infra.Logger, infra.Cache)
	if err != nil {
		infra.Logger.Panicf("error init user repository, %v", err)
	}
	otpRepository, err := otpRepository.NewOTPRepository(infra.Logger, infra.Cache)
	if err != nil {
		infra.Logger.Panicf("error init user repository, %v", err)
	}

	module := &Module{
		Infra:          infra,
		UserRepository: userRepository,
		OTPRepository:  otpRepository,
		AuthRepository: authRepository,
	}
	return module
}
