package model

import (
	otpType "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

const VERIFICATION_PERIOD_IN_SECOND = 300

type Verification struct {
	UID              string
	Receiver         string
	DurationInSecond int
}

func NewVerification(req otpType.CreateVerification) (domain *Verification, err error) {
	domain = &Verification{
		UID:              utils.GenerateUUID(),
		Receiver:         req.Receiver,
		DurationInSecond: VERIFICATION_PERIOD_IN_SECOND,
	}
	return
}
