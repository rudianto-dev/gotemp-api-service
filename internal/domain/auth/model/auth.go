package model

import (
	authType "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model/type"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Username      string
	PlainPassword string
	HashPassword  string
}

func New(req authType.Credential) (domain *Auth, err error) {
	domain = &Auth{
		Username:      req.Username,
		PlainPassword: req.Password,
	}
	if req.Password != "" {
		var hash string
		hash, err = HashPassword(req.Password)
		if err != nil {
			return
		}
		domain.HashPassword = hash
	}
	return
}

func HashPassword(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(bytes), err
}
