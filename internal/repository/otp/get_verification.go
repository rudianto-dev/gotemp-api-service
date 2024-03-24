package otp

import (
	"context"
	"strings"

	"github.com/go-redis/redis"
	otpDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/model"
	otpRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *OTPRepository) GetVerification(ctx context.Context, verificationID string) (otp *otpDomain.Verification, err error) {
	var payload string
	payload, err = s.redis.Get(otpRepository.GetVerificationKey(verificationID)).Result()
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		if err == redis.Nil {
			err = utils.ErrExpiredVerificationOTP
		} else {
			err = utils.ErrRepositoryOTP
		}
		return
	}
	if (payload == "") || (strings.Compare(payload, "") == 0) {
		err = utils.ErrExpiredVerificationOTP
	}
	otp, err = otpRepository.ToVerificationDomain(payload)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrRepositoryOTP
		return
	}
	return
}
