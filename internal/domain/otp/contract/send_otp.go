package otp

import (
	otpType "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/model/type"
)

type SendOTPRequest struct {
	Receiver    string              `json:"receiver" validate:"required"`
	ChannelType otpType.ChannelType `json:"channel_type" validate:"required"`
}

type SendOTPResponse struct {
	ExpiredAt     int64 `json:"expired_at"`
	NextAttemptAt int64 `json:"next_attempt_at"`
}
