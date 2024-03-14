package user

import (
	"context"
	"fmt"

	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) GetByID(ctx context.Context, id string) (*userDomain.User, error) {
	userEntity := UserEntity{}
	query := `
		SELECT id, name, created_at, updated_at
		FROM %s
		WHERE id=:id
	`
	query = fmt.Sprintf(query, sqlUserTable)
	params := map[string]interface{}{
		"id": id,
	}
	rows, err := s.db.Read(ctx, query, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		return toUserDomain(&userEntity), utils.ErrQueryRead
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.StructScan(&userEntity)
		if err != nil {
			s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
			return toUserDomain(&userEntity), utils.ErrQueryRead
		}
	}
	if userEntity.ID == "" {
		return toUserDomain(&userEntity), utils.ErrNotFound
	}
	return toUserDomain(&userEntity), nil
}
