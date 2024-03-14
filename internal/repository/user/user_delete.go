package user

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) Delete(ctx context.Context, tx *sqlx.Tx, user *userDomain.User) (*userDomain.User, error) {
	query := `
		DELETE FROM %s
		WHERE id=:id
	`
	query = fmt.Sprintf(query, sqlUserTable)
	params := map[string]interface{}{
		"id": user.ID,
	}

	err := s.db.Write(ctx, tx, query, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		return user, utils.ErrQueryTxDelete
	}
	return user, nil
}
