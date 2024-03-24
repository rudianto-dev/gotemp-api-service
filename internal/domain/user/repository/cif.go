package repository

import (
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
)

type UserCIFEntity struct {
	ID          string `json:"id" db:"id"`
	UserID      string `json:"user_id" db:"user_id"`
	ReferenceID string `json:"reference_id" db:"reference_id"`
	CreatedAt   int64  `json:"created_at" db:"created_at"`
	UpdatedAt   int64  `json:"updated_at" db:"updated_at"`
}

func ToUserCIFEntity(domain *userDomain.CIF) *UserCIFEntity {
	user := &UserCIFEntity{
		ID:          domain.ID,
		UserID:      domain.UserID,
		ReferenceID: domain.ReferenceID,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
	return user
}

func ToUserCIFDomain(entity *UserCIFEntity) *userDomain.CIF {
	return &userDomain.CIF{
		ID:          entity.ID,
		UserID:      entity.UserID,
		ReferenceID: entity.ReferenceID,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func ToUserCIFDomains(entities []*UserCIFEntity) []*userDomain.CIF {
	domains := []*userDomain.CIF{}
	for _, entity := range entities {
		domains = append(domains, ToUserCIFDomain(entity))
	}
	return domains
}
