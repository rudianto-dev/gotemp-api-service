package client

import (
	"context"
	"fmt"

	clientDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/model"
	clientRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *ClientRepository) List(ctx context.Context) (clients []*clientDomain.Client, err error) {
	clientEntities := []*clientRepository.ClientEntity{}
	params := map[string]interface{}{}
	sqlCommand := `
		SELECT id, client_id, client_secret, status, expired_at
		FROM %s
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlClientTable)
	rows, err := s.db.Read(ctx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryRead
		return
	}
	defer rows.Close()

	for rows.Next() {
		clientEntity := clientRepository.ClientEntity{}
		err = rows.StructScan(&clientEntity)
		if err != nil {
			s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
			err = utils.ErrQueryRead
			return
		}
		clientEntities = append(clientEntities, &clientEntity)
	}
	clients = clientRepository.ToClientDomains(clientEntities)
	return
}
