package client

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	clientRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *ClientRepository) Insert(ctx context.Context, tx *sqlx.Tx, clientEntity *clientRepository.ClientEntity) (err error) {
	sqlCommand := `
		INSERT INTO %s (id, client_id, client_secret, status, expired_at)
		VALUES (:id, :client_id, :client_secret, :status, :expired_at)
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlClientTable)
	params := map[string]interface{}{
		"id":            clientEntity.ID,
		"client_id":     clientEntity.ClientID,
		"client_secret": clientEntity.ClientSecret,
		"status":        clientEntity.Status,
		"expired_at":    clientEntity.ExpiredAt,
	}

	err = s.db.Write(ctx, tx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryTxInsert
		return
	}
	return
}
