package model

import (
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	otpType "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

const OTP_PERIOD_IN_SECOND = 300
const OTP_RESEND_INTERVAL_IN_SECOND = 300

type OTP struct {
	UID              string
	Receiver         string
	ChannelType      otpType.ChannelType
	Code             string
	Secret           string
	DurationInSecond int
	IntervalInSecond int
}

func New(req otpType.CreateOTP) (domain *OTP, err error) {
	UID := utils.GenerateUUID()
	// validate receiver based on channel type
	domain = &OTP{
		UID:              UID,
		Receiver:         req.Receiver,
		ChannelType:      req.ChannelType,
		DurationInSecond: OTP_PERIOD_IN_SECOND,
		IntervalInSecond: OTP_RESEND_INTERVAL_IN_SECOND,
	}
	return
}

func (domain *OTP) Generate(executedAt time.Time) (err error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "local.sig.co",
		AccountName: domain.Receiver,
		Period:      uint(OTP_PERIOD_IN_SECOND),
		Digits:      otp.DigitsSix,
	})
	if err != nil {
		return
	}
	code, err := totp.GenerateCodeCustom(key.Secret(), executedAt, totp.ValidateOpts{
		Period:    uint(OTP_PERIOD_IN_SECOND),
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	if err != nil {
		return
	}
	domain.Secret = key.Secret()
	domain.Code = code

	if !domain.Validate() {
		err = utils.ErrGenerateOTP
		return
	}
	return
}

func (domain *OTP) Validate() (validate bool) {
	validate, _ = totp.ValidateCustom(domain.Code,
		domain.Secret,
		time.Now().UTC(),
		totp.ValidateOpts{
			Period:    uint(OTP_PERIOD_IN_SECOND),
			Skew:      1,
			Digits:    otp.DigitsSix,
			Algorithm: otp.AlgorithmSHA1,
		},
	)
	return
}
