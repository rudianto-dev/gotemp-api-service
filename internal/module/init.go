package module

import (
	"github.com/rudianto-dev/gotemp-api-service/cmd/configuration"
	authInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth"
	otpInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp"
	userInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	utilInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/util"
	authHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/auth"
	otpHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/otp"
	userHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/user"
	utilHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/util"
	authRepository "github.com/rudianto-dev/gotemp-api-service/internal/repository/auth"
	otpRepository "github.com/rudianto-dev/gotemp-api-service/internal/repository/otp"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/repository/user"
	authUseCase "github.com/rudianto-dev/gotemp-api-service/internal/usecase/auth"
	otpUseCase "github.com/rudianto-dev/gotemp-api-service/internal/usecase/otp"
	userUseCase "github.com/rudianto-dev/gotemp-api-service/internal/usecase/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/cache"
	"github.com/rudianto-dev/gotemp-sdk/pkg/database"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
	"github.com/rudianto-dev/gotemp-sdk/pkg/token"
)

type Module struct {
	Infra       *Service
	UtilHandler utilInterface.HandlerAPI
	UserHandler userInterface.HandlerAPI
	OTPHandler  otpInterface.HandlerAPI
	AuthHandler authInterface.HandlerAPI
}

type Service struct {
	Config *configuration.ConfigurationSchema
	Logger *logger.Logger
	Cache  *cache.DataSource
	DB     *database.DB
	JWT    *token.JWT
}

func NewModule(infra *Service) *Module {
	// init repository
	userRepository := userRepository.NewUserRepository(infra.Logger, infra.DB)
	authRepository := authRepository.NewAuthRepository(infra.Logger, infra.Cache)
	otpRepository := otpRepository.NewOTPRepository(infra.Logger, infra.Cache)
	// init use-cases
	authUseCase := authUseCase.NewUseCase(infra.Logger, infra.JWT, infra.DB, authRepository, userRepository, otpRepository)
	userUseCase := userUseCase.NewUseCase(infra.Logger, infra.DB, userRepository)
	otpUseCase := otpUseCase.NewUseCase(infra.Logger, otpRepository)
	// init handler
	module := &Module{
		Infra:       infra,
		UtilHandler: utilHandler.NewHandler(infra.Logger),
		UserHandler: userHandler.NewHandler(infra.Logger, userUseCase),
		OTPHandler:  otpHandler.NewHandler(infra.Logger, otpUseCase),
		AuthHandler: authHandler.NewHandler(infra.Logger, authUseCase),
	}
	return module
}
