package user

import (
	"context"

	authDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model"
	authType "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/model/type"
	userContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/contract"
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
	userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
)

func (s *UserUseCase) Create(ctx context.Context, req userContract.CreateRequest) (res *userContract.CreateResponse, err error) {

	newUser, err := userDomain.New(userType.Create{
		Name:   req.Name,
		Status: userType.USER_STATUS_PREREGISTERED,
	})
	if err != nil {
		return
	}

	newAuth, err := authDomain.New(authType.Credential{
		Username: req.Username,
	})
	if err != nil {
		return
	}

	userEntity := userRepository.ToUserEntity(newUser, newAuth)
	err = s.userRepo.Insert(ctx, nil, userEntity)
	if err != nil {
		return
	}
	userResponse := userContract.UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}

	res = &userContract.CreateResponse{User: userResponse}
	return
}
