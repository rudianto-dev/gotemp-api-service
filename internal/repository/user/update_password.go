package user

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) UpdatePassword(ctx context.Context, tx *sqlx.Tx, ID, password string) (err error) {
	sqlCommand := `
		UPDATE %s SET password=:password, updated_at=:updated_at
		WHERE id=:id
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlUserTable)
	params := map[string]interface{}{
		"id":         ID,
		"password":   password,
		"updated_at": time.Now().Unix(),
	}

	err = s.db.Write(ctx, tx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryTxUpdate
		return
	}
	return
}
