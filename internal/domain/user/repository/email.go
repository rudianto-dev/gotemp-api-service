package repository

import (
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
)

type EmailEntity struct {
	ID        string `json:"id" db:"id"`
	UserID    string `json:"user_id" db:"user_id"`
	Email     string `json:"email" db:"email"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
}

func ToEmailEntity(domain *userDomain.Email) *EmailEntity {
	user := &EmailEntity{
		ID:        domain.ID,
		UserID:    domain.UserID,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
	return user
}

func ToEmailDomain(entity *EmailEntity) *userDomain.Email {
	return &userDomain.Email{
		ID:        entity.ID,
		UserID:    entity.UserID,
		Email:     entity.Email,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func ToEmailDomains(entities []*EmailEntity) []*userDomain.Email {
	domains := []*userDomain.Email{}
	for _, entity := range entities {
		domains = append(domains, ToEmailDomain(entity))
	}
	return domains
}
