package user

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) InsertEmail(ctx context.Context, tx *sqlx.Tx, userEmailEntity *userRepository.EmailEntity) (err error) {
	sqlCommand := `
		INSERT INTO %s (id, user_id, email, created_at, updated_at)
		VALUES (:id, :user_id, :email, :created_at, :updated_at)
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlUserEmailTable)
	params := map[string]interface{}{
		"id":         userEmailEntity.ID,
		"user_id":    userEmailEntity.UserID,
		"email":      userEmailEntity.Email,
		"created_at": userEmailEntity.CreatedAt,
		"updated_at": userEmailEntity.UpdatedAt,
	}

	err = s.db.Write(ctx, tx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryTxInsert
		return
	}
	return
}
