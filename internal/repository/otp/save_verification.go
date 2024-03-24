package otp

import (
	"context"

	otpRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *OTPRepository) SaveVerification(ctx context.Context, req *otpRepository.VerificationEntity) (err error) {
	_, err = s.redis.Set(req.Key, req.Payload, req.TTL).Result()
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrRepositoryOTP
		return
	}
	return
}
