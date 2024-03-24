package otp

import (
	otpInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type OTPHandler struct {
	logger      logger.ILogger
	authUseCase otpInterface.UseCase
}

func NewHandler(logger logger.ILogger, authUseCase otpInterface.UseCase) otpInterface.HandlerAPI {
	return &OTPHandler{
		logger:      logger,
		authUseCase: authUseCase,
	}
}
