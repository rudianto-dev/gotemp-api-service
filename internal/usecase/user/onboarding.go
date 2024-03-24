package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	authDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model"
	authType "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model/type"
	userContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/contract"
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
	userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserUseCase) Onboarding(ctx context.Context, req userContract.OnboardingRequest) (res *userContract.OnboardingResponse, err error) {
	userCIF, err := s.userRepo.GetCIFByReferenceID(ctx, req.CIF.ID)
	if err != nil && err != utils.ErrNotFound {
		return
	}
	if userCIF == nil {
		var (
			userEntity        *userRepository.UserEntity
			userCIFEntity     *userRepository.CIFEntity
			phoneNumberEntity *userRepository.PhoneNumberEntity
		)
		userEntity, userCIFEntity, phoneNumberEntity, err = s.buildRegisterUserOwner(ctx, req.CIF.PhoneNumber, req.CIF.ID)
		if err != nil {
			return
		}
		err = s.RegisterUserOwner(ctx, userEntity, userCIFEntity, phoneNumberEntity)
		if err != nil {
			return
		}
	}
	// UPSERT USER MERCHANT, OUTLET & QRIS
	// put logic here
	res = &userContract.OnboardingResponse{}
	return
}

func (s *UserUseCase) buildRegisterUserOwner(ctx context.Context, phoneNumber, referenceID string) (
	userEntity *userRepository.UserEntity, cifEntity *userRepository.CIFEntity, phoneNumberEntity *userRepository.PhoneNumberEntity, err error) {

	var (
		newUser, user  *userDomain.User
		newAuth        *authDomain.Auth
		newCif         *userDomain.CIF
		newPhoneNumber *userDomain.PhoneNumber
	)
	user, err = s.userRepo.GetByUsername(ctx, phoneNumber)
	if err != nil && err != utils.ErrNotFound {
		return
	}
	if user != nil {
		err = utils.ErrUsernameAlreadyRegistered
		return
	}

	newUser, err = userDomain.New(userType.Create{
		Name:   phoneNumber,
		Status: userType.USER_STATUS_PREREGISTERED,
	})
	if err != nil {
		return
	}
	newAuth = authDomain.New(authType.Credential{Username: phoneNumber})
	newCif, err = userDomain.NewCIF(userType.CreateCIF{
		UserID:      newUser.ID,
		ReferenceID: referenceID,
	})
	if err != nil {
		return
	}
	newPhoneNumber, err = userDomain.NewPhoneNumber(userType.CreatePhoneNumber{
		UserID:      newUser.ID,
		PhoneNumber: phoneNumber,
	})
	if err != nil {
		return
	}
	userEntity = userRepository.ToUserEntity(newUser, newAuth)
	cifEntity = userRepository.ToCIFEntity(newCif)
	phoneNumberEntity = userRepository.ToPhoneNumberEntity(newPhoneNumber)
	return
}

func (s *UserUseCase) RegisterUserOwner(ctx context.Context, userEntity *userRepository.UserEntity,
	userCIFEntity *userRepository.CIFEntity, phoneNumberEntity *userRepository.PhoneNumberEntity) (err error) {

	var tx *sqlx.Tx
	tx, err = s.db.CreateDBTransaction(ctx)
	if err != nil {
		return
	}
	err = s.userRepo.Insert(ctx, tx, userEntity)
	if err != nil {
		tx.Rollback()
		return
	}
	err = s.userRepo.InsertCIF(ctx, tx, userCIFEntity)
	if err != nil {
		tx.Rollback()
		return
	}
	err = s.userRepo.InsertPhoneNumber(ctx, tx, phoneNumberEntity)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
