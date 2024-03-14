package infrastructure

import (
	"github.com/go-redis/redis"
	"github.com/rudianto-dev/gotemp-api-service/cmd/configuration"
	"github.com/rudianto-dev/gotemp-sdk/pkg/database"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	Config *configuration.ConfigurationSchema
	Logger *logger.Logger
	DB     *database.DB
	Redis  *redis.Client
}

func NewInfrastructure(cf *configuration.ConfigurationSchema) *Service {
	// setup logger
	logger := logger.NewLogger(cf.NewLogrus())
	// setup database
	db, err := database.NewDatabase(cf.NewPostgres(), logger, "user", cf.Host.Debug)
	if err != nil {
		log.Panicf("failed connect to database, %v", err)
	}
	db.CheckConnection()
	// setup cache

	return &Service{
		Config: cf,
		Logger: logger,
		DB:     db,
		Redis:  cf.NewRedis(),
	}
}
