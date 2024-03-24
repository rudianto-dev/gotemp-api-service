package auth

import (
	authInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth"
	otpInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp"
	userInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/database"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type AuthUseCase struct {
	logger         logger.ILogger
	db             database.IDatabase
	userRepository userInterface.Repository
	otpRepository  otpInterface.Repository
}

func NewUseCase(logger logger.ILogger, db database.IDatabase, userRepository userInterface.Repository,
	otpRepository otpInterface.Repository) authInterface.UseCase {

	return &AuthUseCase{
		logger:         logger,
		db:             db,
		userRepository: userRepository,
		otpRepository:  otpRepository,
	}
}
