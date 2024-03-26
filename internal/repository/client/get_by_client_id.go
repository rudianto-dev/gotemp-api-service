package client

import (
	"context"
	"fmt"

	clientDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/model"
	clientRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *ClientRepository) GetByClientID(ctx context.Context, clientID string) (client *clientDomain.Client, err error) {
	sqlCommand := `
		SELECT id, client_id, client_secret, status, expired_at
		FROM %s
		WHERE client_id=:client_id
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlClientTable)
	params := map[string]interface{}{
		"client_id": clientID,
	}
	rows, err := s.db.Read(ctx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryRead
		return
	}
	defer rows.Close()

	clientEntity := clientRepository.ClientEntity{}
	if rows.Next() {
		err = rows.StructScan(&clientEntity)
		if err != nil {
			s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
			err = utils.ErrQueryRead
			return
		}
	}
	if clientEntity.ID == "" {
		err = utils.ErrNotFound
		return
	}
	client = clientRepository.ToClientDomain(&clientEntity)
	return
}
