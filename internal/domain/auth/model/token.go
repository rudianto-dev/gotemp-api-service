package model

import (
	authType "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

type Token struct {
	ID        string
	Value     string
	ExpiredAt int64
}

func NewToken(req authType.Token) (domain *Token) {
	domain = &Token{
		ID:        utils.GenerateUUID(),
		Value:     req.Value,
		ExpiredAt: req.ExpiredAt,
	}
	return
}
