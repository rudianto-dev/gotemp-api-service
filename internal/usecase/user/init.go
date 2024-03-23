package user

import (
	userInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type UserUseCase struct {
	logger   logger.ILogger
	userRepo userInterface.Repository
}

func NewUseCase(logger logger.ILogger, userRepo userInterface.Repository) userInterface.UseCase {
	return &UserUseCase{
		logger:   logger,
		userRepo: userRepo,
	}
}
