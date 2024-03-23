package user

import (
	"context"

	userContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/contract"
)

func (s *UserUseCase) Delete(ctx context.Context, req userContract.DeleteRequest) (res *userContract.DeleteResponse, err error) {
	user, err := s.userRepo.GetByID(ctx, req.ID)
	if err != nil {
		return
	}

	err = s.userRepo.Delete(ctx, nil, user.ID)
	if err != nil {
		return
	}
	userResponse := userContract.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	res = &userContract.DeleteResponse{User: userResponse}
	return
}
