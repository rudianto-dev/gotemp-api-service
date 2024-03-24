package user

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) UpdateStatus(ctx context.Context, tx *sqlx.Tx, ID string, status userType.Status) (err error) {
	sqlCommand := `
		UPDATE %s SET status=:status, updated_at=:updated_at
		WHERE id=:id
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlUserTable)
	params := map[string]interface{}{
		"id":         ID,
		"status":     status,
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
