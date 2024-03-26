package client

import (
	"context"

	clientContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/contract"
)

func (s *ClientUseCase) List(ctx context.Context, req clientContract.ListRequest) (res *clientContract.ListResponse, err error) {
	clients, err := s.clientRepo.List(ctx)
	if err != nil {
		return
	}

	clientMap := []clientContract.ClientResponse{}
	for _, client := range clients {
		clientMap = append(clientMap, clientContract.ClientResponse{
			ID:           client.ID,
			ClientID:     client.ClientID,
			ClientSecret: client.ClientSecret,
			Status:       client.Status,
			ExpiredAt:    client.ExpiredAt,
		})
	}
	res = &clientContract.ListResponse{Client: clientMap}
	return
}
