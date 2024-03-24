package auth

import (
	authInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth"
	userInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/database"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type AuthUseCase struct {
	logger   logger.ILogger
	db       database.IDatabase
	userRepo userInterface.Repository
}

func NewUseCase(logger logger.ILogger, db database.IDatabase, userRepo userInterface.Repository) authInterface.UseCase {
	return &AuthUseCase{
		logger:   logger,
		db:       db,
		userRepo: userRepo,
	}
}
