package user

import (
	"context"
	"fmt"

	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) List(ctx context.Context) (users []*userDomain.User, err error) {
	userEntities := []*userRepository.UserEntity{}
	params := map[string]interface{}{}
	sqlCommand := `
		SELECT id, name, status, created_at, updated_at
		FROM %s
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlUserTable)
	rows, err := s.db.Read(ctx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryRead
		return
	}
	defer rows.Close()

	for rows.Next() {
		userEntity := userRepository.UserEntity{}
		err = rows.StructScan(&userEntity)
		if err != nil {
			s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
			err = utils.ErrQueryRead
			return
		}
		userEntities = append(userEntities, &userEntity)
	}
	users = userRepository.ToUserDomains(userEntities)
	return
}
