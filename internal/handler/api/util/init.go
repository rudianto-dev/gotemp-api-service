package util

import (
	utilDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/util"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type UtilHandler struct {
	logger logger.ILogger
}

func NewHandler(logger logger.ILogger) utilDomain.HandlerAPI {
	return &UtilHandler{
		logger: logger,
	}
}
