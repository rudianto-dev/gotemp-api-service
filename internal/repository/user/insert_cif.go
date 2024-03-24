package user

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) InsertCIF(ctx context.Context, tx *sqlx.Tx, userCIFEntity *userRepository.CIFEntity) (err error) {
	sqlCommand := `
		INSERT INTO %s (id, user_id, reference_id, created_at, updated_at)
		VALUES (:id, :user_id, :reference_id, :created_at, :updated_at)
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlUserCifTable)
	params := map[string]interface{}{
		"id":           userCIFEntity.ID,
		"user_id":      userCIFEntity.UserID,
		"reference_id": userCIFEntity.ReferenceID,
		"created_at":   userCIFEntity.CreatedAt,
		"updated_at":   userCIFEntity.UpdatedAt,
	}

	err = s.db.Write(ctx, tx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryTxInsert
		return
	}
	return
}
