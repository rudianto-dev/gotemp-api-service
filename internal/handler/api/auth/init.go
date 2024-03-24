package auth

import (
	authInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type AuthHandler struct {
	logger      logger.ILogger
	authUseCase authInterface.UseCase
}

func NewHandler(logger logger.ILogger, authUseCase authInterface.UseCase) authInterface.HandlerAPI {
	return &AuthHandler{
		logger:      logger,
		authUseCase: authUseCase,
	}
}
