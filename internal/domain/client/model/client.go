package model

import (
	clientType "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

// Client is representing the client data struct
type Client struct {
	ID           string
	ClientID     string
	ClientSecret string
	Status       clientType.Status
	ExpiredAt    int64
}

func New(req clientType.Create) (domain *Client, err error) {
	domain = &Client{
		ID:           utils.GenerateUUID(),
		ClientID:     req.ClientID,
		ClientSecret: req.ClientSecret,
		Status:       req.Status,
		ExpiredAt:    req.ExpiredAt,
	}
	return
}

func Update(req clientType.Edit) (domain *Client, err error) {
	if req.ID == "" {
		err = utils.ErrBadRequest
		return
	}
	domain = &Client{
		ID:           req.ID,
		ClientID:     req.ClientID,
		ClientSecret: req.ClientSecret,
		Status:       req.Status,
		ExpiredAt:    req.ExpiredAt,
	}
	return
}
