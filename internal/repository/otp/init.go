package otp

import (
	"github.com/go-redis/redis"
	otpInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type OTPRepository struct {
	logger *logger.Logger
	redis  *redis.Client
}

func NewOTPRepository(logger *logger.Logger, redis *redis.Client) (otpInterface.Repository, error) {
	otpRepo := &OTPRepository{
		logger: logger,
		redis:  redis,
	}
	return otpRepo, nil
}
