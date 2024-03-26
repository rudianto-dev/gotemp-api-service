package client

import (
	"context"

	clientContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/contract"
)

func (s *ClientUseCase) Delete(ctx context.Context, req clientContract.DeleteRequest) (res *clientContract.DeleteResponse, err error) {
	client, err := s.clientRepo.GetByID(ctx, req.ID)
	if err != nil {
		return
	}

	err = s.clientRepo.Delete(ctx, nil, client.ID)
	if err != nil {
		return
	}
	clientResponse := clientContract.ClientResponse{
		ID:           client.ID,
		ClientID:     client.ClientID,
		ClientSecret: client.ClientSecret,
		Status:       client.Status,
		ExpiredAt:    client.ExpiredAt,
	}
	res = &clientContract.DeleteResponse{Client: clientResponse}
	return
}
