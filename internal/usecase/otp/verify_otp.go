package otp

import (
	"context"

	otpContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/contract"
	otpDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/model"
	otpType "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/model/type"
	otpRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *OTPUseCase) VerifyOTP(ctx context.Context, req otpContract.VerifyOTPRequest) (res *otpContract.VerifyOTPResponse, err error) {
	otp, err := s.otpRepository.Get(ctx, req.Receiver)
	if err != nil {
		return
	}
	otp.Code = req.Code
	if ok := otp.Validate(); !ok {
		err = utils.ErrInvalidOTP
		return
	}
	verification, err := otpDomain.NewVerification(otpType.CreateVerification{
		Receiver: req.Receiver,
	})
	if err != nil {
		return
	}
	verificationEntity, err := otpRepository.ToVerificationEntity(verification)
	if err != nil {
		return
	}
	if err = s.otpRepository.SaveVerification(ctx, verificationEntity); err != nil {
		return
	}
	_ = s.otpRepository.Delete(ctx, req.Receiver)
	res = &otpContract.VerifyOTPResponse{OTPVerificationID: verification.UID}
	return
}
