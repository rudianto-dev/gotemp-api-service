package otp

import (
	alertInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/alert"
	otpInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type OTPUseCase struct {
	logger          logger.ILogger
	otpRepository   otpInterface.Repository
	alertRepository alertInterface.Repository
}

func NewUseCase(logger logger.ILogger, otpRepository otpInterface.Repository, alertRepository alertInterface.Repository) otpInterface.UseCase {
	return &OTPUseCase{
		logger:          logger,
		otpRepository:   otpRepository,
		alertRepository: alertRepository,
	}
}
