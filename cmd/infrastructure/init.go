package infrastructure

import (
	"time"

	"github.com/rudianto-dev/gotemp-api-service/cmd/configuration"
	"github.com/rudianto-dev/gotemp-sdk/pkg/cache"
	"github.com/rudianto-dev/gotemp-sdk/pkg/database"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
	"github.com/rudianto-dev/gotemp-sdk/pkg/token"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	Config *configuration.ConfigurationSchema
	Logger *logger.Logger
	DB     *database.DB
	Cache  *cache.DataSource
	JWT    *token.JWT
}

func New(cf *configuration.ConfigurationSchema) *Service {
	// setup logger
	logger := logger.NewLogger(cf.NewLogrus())
	// setup database
	db, err := database.NewDatabase(cf.NewPostgres(), logger, "user", cf.Host.Debug)
	if err != nil {
		log.Panicf("failed connect to database, %v", err)
	}
	db.CheckConnection()
	// setup cache
	cache, err := cache.NewCache(cf.NewRedis(), logger)
	if err != nil {
		log.Panicf("failed connect to database, %v", err)
	}
	// setup jwt
	jwt := token.New([]byte(cf.JWT.Private), []byte(cf.JWT.Public), time.Hour*time.Duration(cf.JWT.ExpireInHour), logger)

	return &Service{
		Config: cf,
		Logger: logger,
		DB:     db,
		Cache:  cache,
		JWT:    jwt,
	}
}
