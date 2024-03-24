package model

import (
	otpType "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

type Verification struct {
	UID              string
	Receiver         string
	DurationInSecond int
}

func NewVerification(req otpType.CreateVerification) (domain *Verification, err error) {
	domain = &Verification{
		UID:      utils.GenerateUUID(),
		Receiver: req.Receiver,
	}
	return
}
