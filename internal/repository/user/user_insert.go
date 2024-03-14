package user

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) Create(ctx context.Context, tx *sqlx.Tx, userDomain *userDomain.User) (*userDomain.User, error) {
	user := toUserEntity(userDomain)
	query := `
		INSERT INTO %s (id, name, created_at, updated_at)
		VALUES (:id, :name, :created_at, :updated_at)
	`
	query = fmt.Sprintf(query, sqlUserTable)
	params := map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	}

	err := s.db.Write(ctx, tx, query, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		return userDomain, utils.ErrQueryTxInsert
	}
	return userDomain, nil
}
