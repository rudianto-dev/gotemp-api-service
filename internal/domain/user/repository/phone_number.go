package repository

import (
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
)

type PhoneNumberEntity struct {
	ID          string `json:"id" db:"id"`
	UserID      string `json:"user_id" db:"user_id"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	CreatedAt   int64  `json:"created_at" db:"created_at"`
	UpdatedAt   int64  `json:"updated_at" db:"updated_at"`
}

func ToPhoneNumberEntity(domain *userDomain.PhoneNumber) *PhoneNumberEntity {
	user := &PhoneNumberEntity{
		ID:          domain.ID,
		UserID:      domain.UserID,
		PhoneNumber: domain.PhoneNumber,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
	return user
}

func ToPhoneNumberDomain(entity *PhoneNumberEntity) *userDomain.PhoneNumber {
	return &userDomain.PhoneNumber{
		ID:          entity.ID,
		UserID:      entity.UserID,
		PhoneNumber: entity.PhoneNumber,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func ToPhoneNumberDomains(entities []*PhoneNumberEntity) []*userDomain.PhoneNumber {
	domains := []*userDomain.PhoneNumber{}
	for _, entity := range entities {
		domains = append(domains, ToPhoneNumberDomain(entity))
	}
	return domains
}
