package user

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) InsertDeviceInstance(ctx context.Context, tx *sqlx.Tx, userDeviceInstanceEntity *userRepository.DeviceInstanceEntity) (err error) {
	sqlCommand := `
		INSERT INTO %s (id, user_id, device_id, instance_id, created_at, updated_at)
		VALUES (:id, :user_id, :device_id, :instance_id, :created_at, :updated_at)
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlUserDeviceInstanceTable)
	params := map[string]interface{}{
		"id":          userDeviceInstanceEntity.ID,
		"user_id":     userDeviceInstanceEntity.UserID,
		"device_id":   userDeviceInstanceEntity.DeviceID,
		"instance_id": userDeviceInstanceEntity.InstanceID,
		"created_at":  userDeviceInstanceEntity.CreatedAt,
		"updated_at":  userDeviceInstanceEntity.UpdatedAt,
	}

	err = s.db.Write(ctx, tx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryTxInsert
		return
	}
	return
}
