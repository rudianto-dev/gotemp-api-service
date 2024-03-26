package client

import (
	clientInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/client"
	"github.com/rudianto-dev/gotemp-sdk/pkg/cache"
	"github.com/rudianto-dev/gotemp-sdk/pkg/database"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type ClientRepository struct {
	logger *logger.Logger
	db     *database.DB
	cache  *cache.DataSource
}

const (
	sqlClientTable = "clients"
)

func NewClientRepository(logger *logger.Logger, db *database.DB, cache *cache.DataSource) clientInterface.Repository {
	return &ClientRepository{
		logger: logger,
		db:     db,
		cache:  cache,
	}
}
