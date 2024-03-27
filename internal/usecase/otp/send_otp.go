package otp

import (
	"context"
	"fmt"
	"time"

	alertDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/alert/model"
	alertType "github.com/rudianto-dev/gotemp-api-service/internal/domain/alert/model/type"
	alertRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/alert/repository"
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

	alert := alertDomain.New(alertType.Create{
		Message:  fmt.Sprintf("Kode OTP %s. Untuk keamanan, jangan bagikan kode ini ke orang lain.", otp.Code),
		Sender:   "OCBC merchant",
		Receiver: req.Receiver,
		Channel:  alertType.ALERT_CHANNEL_TELEGRAM,
	})
	go func() {
		if err := s.alertRepository.Send(context.Background(), alertRepository.ToAlertEntity(alert)); err != nil {
			s.logger.Error(err)
		}
	}()

	res = &otpContract.SendOTPResponse{
		ExpiredAt:     now.Add(time.Duration(otp.DurationInSecond) * time.Second).Unix(),
		NextAttemptAt: now.Add(time.Duration(otp.IntervalInSecond) * time.Second).Unix(),
	}
	return
}
