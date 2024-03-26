package client

import clientType "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/model/type"

type ClientResponse struct {
	ID           string            `json:"id" validate:"required"`
	ClientID     string            `json:"client_id" validate:"required"`
	ClientSecret string            `json:"client_secret" validate:"required"`
	Status       clientType.Status `json:"status" validate:"required"`
	ExpiredAt    int64             `json:"expired_at" validate:"required"`
}

type CreateRequest struct {
	ClientID     string            `json:"client_id" validate:"required"`
	ClientSecret string            `json:"client_secret" validate:"required"`
	Status       clientType.Status `json:"status" validate:"required"`
	ExpiredAt    int64             `json:"expired_at" validate:"required"`
}

type CreateResponse struct {
	Client ClientResponse `json:"client"`
}
