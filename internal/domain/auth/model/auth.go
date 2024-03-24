package model

import (
	authType "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model/type"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Username string
	Password string
}

func New(req authType.Credential) (domain *Auth) {
	domain = &Auth{
		Username: req.Username,
		Password: req.Password,
	}
	return
}

func (domain *Auth) HashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(domain.Password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (domain *Auth) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(domain.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}