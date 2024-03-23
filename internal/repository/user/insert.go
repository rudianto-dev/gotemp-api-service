package user

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) Insert(ctx context.Context, tx *sqlx.Tx, userEntity *userRepository.UserEntity) (err error) {
	sqlCommand := `
		INSERT INTO %s (id, name, username, status, created_at, updated_at)
		VALUES (:id, :name, :username, :status, :created_at, :updated_at)
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlUserTable)
	params := map[string]interface{}{
		"id":         userEntity.ID,
		"name":       userEntity.Name,
		"username":   userEntity.Username,
		"status":     userEntity.Status,
		"created_at": userEntity.CreatedAt,
		"updated_at": userEntity.UpdatedAt,
	}

	err = s.db.Write(ctx, tx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryTxInsert
		return
	}
	return
}
