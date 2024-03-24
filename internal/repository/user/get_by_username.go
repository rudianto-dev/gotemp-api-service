package user

import (
	"context"
	"fmt"

	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) GetByUsername(ctx context.Context, username string) (user *userDomain.User, err error) {
	sqlCommand := `
		SELECT id, name, username, password, status, role_type, created_at, updated_at
		FROM %s
		WHERE username=:username
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlUserTable)
	params := map[string]interface{}{
		"username": username,
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
