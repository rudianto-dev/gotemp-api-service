package user

import (
	"context"
	"fmt"

	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserRepository) ListPhoneNumberGetByUserID(ctx context.Context, UserID string) (users []*userDomain.PhoneNumber, err error) {
	userEntities := []*userRepository.PhoneNumberEntity{}
	sqlCommand := `
		SELECT id, user_id, phone_number, created_at, updated_at
		FROM %s
		WHERE user_id=:user_id
	`
	sqlCommand = fmt.Sprintf(sqlCommand, sqlUserPhoneNumberTable)
	params := map[string]interface{}{
		"user_id": UserID,
	}
	rows, err := s.db.Read(ctx, sqlCommand, params)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
		err = utils.ErrQueryRead
		return
	}
	defer rows.Close()

	for rows.Next() {
		phoneNumberEntity := userRepository.PhoneNumberEntity{}
		err = rows.StructScan(&phoneNumberEntity)
		if err != nil {
			s.logger.ErrorWithContext(ctx, utils.ERROR_REPOSITORY_STAGE, err.Error())
			err = utils.ErrQueryRead
			return
		}
		userEntities = append(userEntities, &phoneNumberEntity)
	}
	users = userRepository.ToPhoneNumberDomains(userEntities)
	return
}
