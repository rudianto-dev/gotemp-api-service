package repository

import (
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
)

type CIFEntity struct {
	ID          string `json:"id" db:"id"`
	UserID      string `json:"user_id" db:"user_id"`
	ReferenceID string `json:"reference_id" db:"reference_id"`
	CreatedAt   int64  `json:"created_at" db:"created_at"`
	UpdatedAt   int64  `json:"updated_at" db:"updated_at"`
}

func ToCIFEntity(domain *userDomain.CIF) *CIFEntity {
	user := &CIFEntity{
		ID:          domain.ID,
		UserID:      domain.UserID,
		ReferenceID: domain.ReferenceID,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
	return user
}

func ToCIFDomain(entity *CIFEntity) *userDomain.CIF {
	return &userDomain.CIF{
		ID:          entity.ID,
		UserID:      entity.UserID,
		ReferenceID: entity.ReferenceID,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func ToCIFDomains(entities []*CIFEntity) []*userDomain.CIF {
	domains := []*userDomain.CIF{}
	for _, entity := range entities {
		domains = append(domains, ToCIFDomain(entity))
	}
	return domains
}
