package repository

import (
	"encoding/json"
	"fmt"
	"time"

	otpDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/otp/model"
	"github.com/spf13/cast"
)

const (
	prefixVerification = "verification_otp:%s"
)

type VerificationEntity struct {
	Key     string
	Payload string
	TTL     time.Duration
}

func ToVerificationEntity(domain *otpDomain.Verification) (entity *VerificationEntity, err error) {
	payload, err := json.Marshal(domain)
	if err != nil {
		return
	}
	entity = &VerificationEntity{
		Key:     fmt.Sprintf(prefixVerification, domain.UID),
		Payload: cast.ToString(payload),
		TTL:     time.Second * time.Duration(domain.DurationInSecond),
	}
	return
}

func ToVerificationDomain(payload string) (domain *otpDomain.Verification, err error) {
	res := &otpDomain.Verification{}
	err = json.Unmarshal([]byte(payload), res)
	if err != nil {
		return
	}
	domain = res
	return
}

func GetVerificationKey(receiver string) string {
	return fmt.Sprintf(prefixVerification, receiver)
}
