package repository

import (
	"encoding/json"
	"fmt"
	"time"

	authDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model"
	"github.com/spf13/cast"
)

const (
	prefixToken = "token:%s"
)

const REFRESH_PERIOD_IN_HOUR = 5

type TokenEntity struct {
	Key     string
	Payload string
	TTL     time.Duration
}

func ToTokenEntity(domain *authDomain.Token) (entity *TokenEntity, err error) {
	payload, err := json.Marshal(domain)
	if err != nil {
		return
	}
	entity = &TokenEntity{
		Key:     fmt.Sprintf(prefixToken, domain.ID),
		Payload: cast.ToString(payload),
		TTL:     time.Hour * time.Duration(REFRESH_PERIOD_IN_HOUR),
	}
	return
}

func ToTokenDomain(payload string) (domain *authDomain.Token, err error) {
	res := &authDomain.Token{}
	err = json.Unmarshal([]byte(payload), res)
	if err != nil {
		return
	}
	domain = res
	return
}

func GetTokenKey(receiver string) string {
	return fmt.Sprintf(prefixToken, receiver)
}
