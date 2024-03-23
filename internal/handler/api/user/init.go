package user

import (
	userInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type UserHandler struct {
	logger      logger.ILogger
	userUseCase userInterface.UseCase
}

func NewHandler(logger logger.ILogger, userUseCase userInterface.UseCase) userInterface.HandlerAPI {
	return &UserHandler{
		logger:      logger,
		userUseCase: userUseCase,
	}
}
