package repository

import (
	"encoding/json"
	"fmt"
	"time"

	clientDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/model"
	"github.com/spf13/cast"
)

const (
	prefixOTP = "otp:%s"
)

type CacheEntity struct {
	Key     string
	Payload string
	TTL     time.Duration
}

func ToCacheEntity(domain *clientDomain.Client) (entity *CacheEntity, err error) {
	payload, err := json.Marshal(domain)
	if err != nil {
		return
	}
	entity = &CacheEntity{
		Key:     fmt.Sprintf(prefixOTP, domain.ClientID),
		Payload: cast.ToString(payload),
		TTL:     -1,
	}
	return
}

func ToOTPDomain(payload string) (domain *clientDomain.Client, err error) {
	res := &clientDomain.Client{}
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
