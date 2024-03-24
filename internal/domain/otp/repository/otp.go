package repository

import (
	"encoding/json"
	"fmt"
	"time"

	otpDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/model"
	"github.com/spf13/cast"
)

const (
	prefixOTP = "otp:%s"
)

type OTPEntity struct {
	Key     string
	Payload string
	TTL     time.Duration
}

func ToOTPEntity(domain *otpDomain.OTP) (entity *OTPEntity, err error) {
	payload, err := json.Marshal(domain)
	if err != nil {
		return
	}
	entity = &OTPEntity{
		Key:     fmt.Sprintf(prefixOTP, domain.Receiver),
		Payload: cast.ToString(payload),
		TTL:     time.Second * time.Duration(domain.DurationInSecond),
	}
	return
}

func ToOTPDomain(payload string) (domain *otpDomain.OTP, err error) {
	res := &otpDomain.OTP{}
	err = json.Unmarshal([]byte(payload), res)
	if err != nil {
		return
	}
	domain = res
	return
}

func GetOTPKey(receiver string) string {
	return fmt.Sprintf(prefixOTP, receiver)
}
