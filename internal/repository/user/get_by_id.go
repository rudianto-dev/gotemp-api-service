package user

import (
	"context"
	"fmt"

	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) GetByID(ctx context.Context, ID string) (user *userDomain.User, err error) {
	sqlCommand := `
		SELECT id, name, created_at, updated_at
		FROM %s
		WHERE id=:id
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlUserTable)
	params := map[string]interface{}{
		"id": ID,
	}
	rows, err := s.db.Read(ctx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryRead
		return
	}
	defer rows.Close()

	userEntity := userRepository.UserEntity{}
	if rows.Next() {
		err = rows.StructScan(&userEntity)
		if err != nil {
			s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
			err = utils.ErrQueryRead
			return
		}
	}
	if userEntity.ID == "" {
		err = utils.ErrNotFound
		return
	}
	user = userRepository.ToUserDomain(&userEntity)
	return
}
