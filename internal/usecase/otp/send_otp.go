package otp

import (
	"context"
	"time"

	otpContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/contract"
	otpDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/model"
	otpType "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/model/type"
	otpRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/repository"
)

func (s *OTPUseCase) SendOTP(ctx context.Context, req otpContract.SendOTPRequest) (res *otpContract.SendOTPResponse, err error) {
	otp, err := otpDomain.New(otpType.CreateOTP{
		Receiver:    req.Receiver,
		ChannelType: req.ChannelType,
	})
	if err != nil {
		return
	}
	now := time.Now()
	err = otp.Generate(now)
	if err != nil {
		return
	}
	otpEntity, err := otpRepository.ToOTPEntity(otp)
	if err != nil {
		return
	}
	err = s.otpRepository.Save(ctx, otpEntity)
	res = &otpContract.SendOTPResponse{
		ExpiredAt:     now.Add(time.Duration(otp.DurationInSecond) * time.Second).Unix(),
		NextAttemptAt: now.Add(time.Duration(otp.IntervalInSecond) * time.Second).Unix(),
	}
	return
}
