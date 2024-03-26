package client

import (
	clientInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/client"
	"github.com/rudianto-dev/gotemp-sdk/pkg/database"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type ClientUseCase struct {
	logger     logger.ILogger
	db         database.IDatabase
	clientRepo clientInterface.Repository
}

func NewUseCase(logger logger.ILogger, db database.IDatabase, clientRepo clientInterface.Repository) clientInterface.UseCase {
	return &ClientUseCase{
		logger:     logger,
		db:         db,
		clientRepo: clientRepo,
	}
}
