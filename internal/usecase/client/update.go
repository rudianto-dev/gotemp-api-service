package client

import (
	"context"

	clientContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/contract"
	clientRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/repository"
)

func (s *ClientUseCase) Update(ctx context.Context, req clientContract.UpdateRequest) (res *clientContract.UpdateResponse, err error) {
	updateClient, err := s.clientRepo.GetByID(ctx, req.ID)
	if err != nil {
		return
	}
	oldClientID := updateClient.ClientID
	updateClient.ClientID = req.ClientID
	updateClient.ClientSecret = req.ClientSecret
	updateClient.Status = req.Status
	updateClient.ExpiredAt = req.ExpiredAt

	clientEntity := clientRepository.ToClientEntity(updateClient)
	err = s.clientRepo.Update(ctx, nil, clientEntity)
	if err != nil {
		return
	}

	cache, err := clientRepository.ToCacheEntity(updateClient)
	if err != nil {
		return
	}
	_ = s.clientRepo.DeleteCache(ctx, oldClientID)
	err = s.clientRepo.SaveCache(ctx, cache)
	if err != nil {
		return
	}

	clientResponse := clientContract.ClientResponse{
		ID:           updateClient.ID,
		ClientID:     updateClient.ClientID,
		ClientSecret: updateClient.ClientSecret,
		Status:       updateClient.Status,
		ExpiredAt:    updateClient.ExpiredAt,
	}
	res = &clientContract.UpdateResponse{Client: clientResponse}
	return
}
