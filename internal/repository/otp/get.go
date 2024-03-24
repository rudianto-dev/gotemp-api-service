package otp

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis"
	otpDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/model"
	otpRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *OTPRepository) Get(ctx context.Context, receiver string) (otp *otpDomain.OTP, err error) {
	var payload string
	payload, err = s.redis.Get(otpRepository.GetOTPKey(receiver)).Result()
	fmt.Println("sss ", payload)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		if err == redis.Nil {
			err = utils.ErrExpiredOTP
		} else {
			err = utils.ErrRepositoryOTP
		}
		return
	}
	if (payload == "") || (strings.Compare(payload, "") == 0) {
		err = utils.ErrExpiredOTP
		return
	}
	otp, err = otpRepository.ToOTPDomain(payload)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrRepositoryOTP
		return
	}
	return
}
