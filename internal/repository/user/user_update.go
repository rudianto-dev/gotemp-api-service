package user

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) Update(ctx context.Context, tx *sqlx.Tx, user *userDomain.User) (*userDomain.User, error) {
	query := `
		UPDATE %s SET name=:name, updated_at=:updated_at
		WHERE id=:id
	`
	query = fmt.Sprintf(query, sqlUserTable)
	params := map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"updated_at": user.UpdatedAt,
	}

	err := s.db.Write(ctx, tx, query, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		return user, utils.ErrQueryTxUpdate
	}
	return user, nil
}
