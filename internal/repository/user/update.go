package user

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) Update(ctx context.Context, tx *sqlx.Tx, userEntity *userRepository.UserEntity) (err error) {
	sqlCommand := `
		UPDATE %s SET name=:name, updated_at=:updated_at
		WHERE id=:id
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlUserTable)
	params := map[string]interface{}{
		"id":         userEntity.ID,
		"name":       userEntity.Name,
		"updated_at": userEntity.UpdatedAt,
	}

	err = s.db.Write(ctx, tx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryTxUpdate
		return
	}
	return
}
