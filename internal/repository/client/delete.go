package client

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *ClientRepository) Delete(ctx context.Context, tx *sqlx.Tx, ID string) (err error) {
	sqlCommand := `
		DELETE FROM %s
		WHERE id=:id
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlClientTable)
	params := map[string]interface{}{
		"id": ID,
	}

	err = s.db.Write(ctx, tx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryTxDelete
		return
	}
	return
}
