package user

import (
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type UserHandler struct {
	logger      logger.ILogger
	userUseCase userDomain.IUseCase
}

func NewHandler(logger logger.ILogger, userUseCase userDomain.IUseCase) userDomain.IHandlerAPI {
	return &UserHandler{
		logger:      logger,
		userUseCase: userUseCase,
	}
}
