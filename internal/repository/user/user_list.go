package user

import (
	"context"
	"fmt"

	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) GetList(ctx context.Context) ([]*userDomain.User, error) {
	userEntities := []*UserEntity{}
	params := map[string]interface{}{}
	query := `
		SELECT id, name, created_at, updated_at
		FROM %s
	`
	query = fmt.Sprintf(query, sqlUserTable)
	rows, err := s.db.Read(ctx, query, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		return toUserDomains(userEntities), utils.ErrQueryRead
	}
	defer rows.Close()
	for rows.Next() {
		userEntity := UserEntity{}
		err = rows.StructScan(&userEntity)
		if err != nil {
			s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
			return toUserDomains(userEntities), utils.ErrQueryRead
		}

		userEntities = append(userEntities, &userEntity)
	}
	return toUserDomains(userEntities), nil
}
