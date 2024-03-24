package user

import (
	"context"
	"fmt"

	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) GetCIFByReferenceID(ctx context.Context, referenceID string) (CIF *userDomain.CIF, err error) {
	sqlCommand := `
		SELECT id, user_id, reference_id, created_at, updated_at
		FROM %s
		WHERE reference_id=:reference_id
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlUserCifTable)
	params := map[string]interface{}{
		"reference_id": referenceID,
	}
	rows, err := s.db.Read(ctx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryRead
		return
	}
	defer rows.Close()

	UserCIFEntity := userRepository.UserCIFEntity{}
	if rows.Next() {
		err = rows.StructScan(&UserCIFEntity)
		if err != nil {
			s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
			err = utils.ErrQueryRead
			return
		}
	}
	if UserCIFEntity.ID == "" {
		err = utils.ErrNotFound
		return
	}
	CIF = userRepository.ToUserCIFDomain(&UserCIFEntity)
	return
}
