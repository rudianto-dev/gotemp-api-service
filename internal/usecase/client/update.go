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
	updateClient.ClientID = req.ClientID
	updateClient.ClientSecret = req.ClientSecret
	updateClient.Status = req.Status
	updateClient.ExpiredAt = req.ExpiredAt

	userEntity := clientRepository.ToClientEntity(updateClient)
	err = s.clientRepo.Update(ctx, nil, userEntity)
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
