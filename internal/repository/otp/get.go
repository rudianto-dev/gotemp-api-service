package otp

import (
	"context"
	"strings"

	otpDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/model"
	otpRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *OTPRepository) Get(ctx context.Context, receiver string) (otp *otpDomain.OTP, err error) {
	var payload string
	payload, err = s.cache.Get(ctx, otpRepository.GetOTPKey(receiver))
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrRepositoryOTP
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
