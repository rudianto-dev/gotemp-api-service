package client

import (
	"context"

	clientContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/contract"
	clientDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/model"
	clientType "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/model/type"
	clientRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *ClientUseCase) Create(ctx context.Context, req clientContract.CreateRequest) (res *clientContract.CreateResponse, err error) {
	client, err := s.clientRepo.GetByClientID(ctx, req.ClientID)
	if err != nil && err != utils.ErrNotFound {
		return
	}
	if client != nil {
		err = utils.ErrClientIDAlreadyRegistered
		return
	}
	newClient, err := clientDomain.New(clientType.Create{
		ClientID:     req.ClientID,
		ClientSecret: req.ClientSecret,
		Status:       clientType.Status(req.Status),
		ExpiredAt:    req.ExpiredAt,
	})
	if err != nil {
		return
	}
	clientEntity := clientRepository.ToClientEntity(newClient)
	err = s.clientRepo.Insert(ctx, nil, clientEntity)
	if err != nil {
		return
	}

	cache, err := clientRepository.ToCacheEntity(newClient)
	if err != nil {
		return
	}
	err = s.clientRepo.SaveCache(ctx, cache)
	if err != nil {
		return
	}

	clientResponse := clientContract.ClientResponse{
		ID:           newClient.ID,
		ClientID:     newClient.ClientID,
		ClientSecret: newClient.ClientSecret,
		Status:       newClient.Status,
		ExpiredAt:    newClient.ExpiredAt,
	}
	res = &clientContract.CreateResponse{Client: clientResponse}
	return
}
