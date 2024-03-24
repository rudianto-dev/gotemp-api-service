package otp

import (
	"context"
	"net/http"

	otpContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/contract"
	otpDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/model"
	otpRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/repository"
)

type Repository interface {
	Save(ctx context.Context, req *otpRepository.OTPEntity) error
	Get(ctx context.Context, receiver string) (*otpDomain.OTP, error)
	Delete(ctx context.Context, receiver string) error

	SaveVerification(ctx context.Context, req *otpRepository.VerificationEntity) error
	GetVerification(ctx context.Context, verificationID string) (*otpDomain.Verification, error)
	DeleteVerification(ctx context.Context, verificationID string) error
}

type UseCase interface {
	SendOTP(ctx context.Context, req otpContract.SendOTPRequest) (*otpContract.SendOTPResponse, error)
	VerifyOTP(ctx context.Context, req otpContract.VerifyOTPRequest) (*otpContract.VerifyOTPResponse, error)
}

type HandlerAPI interface {
	SendOTP(w http.ResponseWriter, r *http.Request)
	VerifyOTP(w http.ResponseWriter, r *http.Request)
}
