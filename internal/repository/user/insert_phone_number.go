package user

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) InsertPhoneNumber(ctx context.Context, tx *sqlx.Tx, userPhoneNumberEntity *userRepository.PhoneNumberEntity) (err error) {
	sqlCommand := `
		INSERT INTO %s (id, user_id, phone_number, created_at, updated_at)
		VALUES (:id, :user_id, :phone_number, :created_at, :updated_at)
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlUserPhoneNumberTable)
	params := map[string]interface{}{
		"id":           userPhoneNumberEntity.ID,
		"user_id":      userPhoneNumberEntity.UserID,
		"phone_number": userPhoneNumberEntity.PhoneNumber,
		"created_at":   userPhoneNumberEntity.CreatedAt,
		"updated_at":   userPhoneNumberEntity.UpdatedAt,
	}

	err = s.db.Write(ctx, tx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryTxInsert
		return
	}
	return
}
