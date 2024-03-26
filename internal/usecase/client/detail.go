package client

import (
	"context"

	clientContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/contract"
)

func (s *ClientUseCase) Detail(ctx context.Context, req clientContract.DetailRequest) (res *clientContract.DetailResponse, err error) {
	client, err := s.clientRepo.GetByID(ctx, req.ID)
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
	res = &clientContract.DetailResponse{Client: clientResponse}
	return
}
