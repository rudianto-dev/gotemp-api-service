package user

import (
	"context"

	userContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/contract"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
)

func (s *UserUseCase) Update(ctx context.Context, req userContract.UpdateRequest) (res *userContract.UpdateResponse, err error) {
	updateUser, err := s.userRepo.GetByID(ctx, req.ID)
	if err != nil {
		return
	}
	updateUser.Name = req.Name

	userEntity := userRepository.ToUserEntity(updateUser, nil)
	err = s.userRepo.Update(ctx, nil, userEntity)
	if err != nil {
		return
	}
	userResponse := userContract.UserResponse{
		ID:        updateUser.ID,
		Name:      updateUser.Name,
		CreatedAt: updateUser.CreatedAt,
		UpdatedAt: updateUser.UpdatedAt,
	}
	res = &userContract.UpdateResponse{User: userResponse}
	return
}
