package user

import (
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type UserUseCase struct {
	logger   logger.ILogger
	userRepo userDomain.IRepository
}

func NewUseCase(logger logger.ILogger, userRepo userDomain.IRepository) userDomain.IUseCase {
	return &UserUseCase{
		logger:   logger,
		userRepo: userRepo,
	}
}
