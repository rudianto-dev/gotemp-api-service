package otp

import (
	"context"

	otpRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *OTPRepository) SaveVerification(ctx context.Context, req *otpRepository.VerificationEntity) (err error) {
	err = s.cache.Save(ctx, req.Key, req.Payload, req.TTL)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrRepositoryOTP
		return
	}
	return
}
