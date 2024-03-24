package otp

import (
	otpInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type OTPUseCase struct {
	logger        logger.ILogger
	otpRepository otpInterface.Repository
}

func NewUseCase(logger logger.ILogger, otpRepository otpInterface.Repository) otpInterface.UseCase {
	return &OTPUseCase{
		logger:        logger,
		otpRepository: otpRepository,
	}
}
