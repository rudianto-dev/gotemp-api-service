package user

import (
	userInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/database"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type UserUseCase struct {
	logger   logger.ILogger
	db       database.IDatabase
	userRepo userInterface.Repository
}

func NewUseCase(logger logger.ILogger, db database.IDatabase, userRepo userInterface.Repository) userInterface.UseCase {
	return &UserUseCase{
		logger:   logger,
		db:       db,
		userRepo: userRepo,
	}
}
