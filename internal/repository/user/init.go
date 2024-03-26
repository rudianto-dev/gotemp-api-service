package user

import (
	userInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/database"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type UserRepository struct {
	logger *logger.Logger
	db     *database.DB
}

const (
	sqlUserTable               = "users"
	sqlUserPhoneNumberTable    = "user_phone_numbers"
	sqlUserEmailTable          = "user_emails"
	sqlUserDeviceInstanceTable = "user_device_instances"
	sqlUserCifTable            = "user_cif"
)

func NewUserRepository(logger *logger.Logger, db *database.DB) userInterface.Repository {
	return &UserRepository{
		logger: logger,
		db:     db,
	}
}
