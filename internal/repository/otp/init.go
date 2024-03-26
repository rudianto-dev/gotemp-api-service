package otp

import (
	otpInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp"
	"github.com/rudianto-dev/gotemp-sdk/pkg/cache"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type OTPRepository struct {
	logger *logger.Logger
	cache  *cache.DataSource
}

func NewOTPRepository(logger *logger.Logger, cache *cache.DataSource) otpInterface.Repository {
	return &OTPRepository{
		logger: logger,
		cache:  cache,
	}
}
