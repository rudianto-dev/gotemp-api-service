package repository

import (
	"fmt"
	"time"

	clientDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/model"
)

const (
	prefixOTP = "client_credential:%s"
)

type CacheEntity struct {
	Key     string
	Payload string
	TTL     time.Duration
}

func ToCacheEntity(domain *clientDomain.Client) (entity *CacheEntity, err error) {
	entity = &CacheEntity{
		Key:     fmt.Sprintf(prefixOTP, domain.ClientID),
		Payload: domain.ClientSecret,
		TTL:     -1,
	}
	return
}

func GetClientKey(receiver string) string {
	return fmt.Sprintf(prefixOTP, receiver)
}
