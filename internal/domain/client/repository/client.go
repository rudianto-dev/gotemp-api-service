package repository

import (
	clientDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/model"
	clientType "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/model/type"
)

type ClientEntity struct {
	ID           string `json:"id" db:"id"`
	ClientID     string `json:"client_id" db:"id"`
	ClientSecret string `json:"client_secret" db:"id"`
	Status       int8   `json:"status" db:"id"`
	ExpiredAt    int64  `json:"expired_at" db:"id"`
}

func ToClientEntity(domain *clientDomain.Client) *ClientEntity {
	user := &ClientEntity{
		ID:           domain.ID,
		ClientID:     domain.ClientID,
		ClientSecret: domain.ClientSecret,
		Status:       int8(domain.Status),
		ExpiredAt:    domain.ExpiredAt,
	}
	return user
}

func ToClientDomain(entity *ClientEntity) *clientDomain.Client {
	return &clientDomain.Client{
		ID:           entity.ID,
		ClientID:     entity.ClientID,
		ClientSecret: entity.ClientSecret,
		Status:       clientType.Status(entity.Status),
		ExpiredAt:    entity.ExpiredAt,
	}
}

func ToClientDomains(entities []*ClientEntity) []*clientDomain.Client {
	domains := []*clientDomain.Client{}
	for _, entity := range entities {
		domains = append(domains, ToClientDomain(entity))
	}
	return domains
}
