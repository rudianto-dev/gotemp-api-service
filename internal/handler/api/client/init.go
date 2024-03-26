package client

import (
	clientInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/client"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type ClientHandler struct {
	logger        logger.ILogger
	clientUseCase clientInterface.UseCase
}

func NewHandler(logger logger.ILogger, clientUseCase clientInterface.UseCase) clientInterface.HandlerAPI {
	return &ClientHandler{
		logger:        logger,
		clientUseCase: clientUseCase,
	}
}
