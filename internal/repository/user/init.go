package user

import (
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/database"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type UserRepository struct {
	logger *logger.Logger
	db     *database.DB
}

const (
	sqlUserTable = "users"
)

func NewUserRepository(logger *logger.Logger, db *database.DB) (userDomain.IRepository, error) {
	userRP := &UserRepository{
		logger: logger,
		db:     db,
	}
	return userRP, nil
}
