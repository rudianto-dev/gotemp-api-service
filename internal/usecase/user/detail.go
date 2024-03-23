package user

import (
	"context"

	userContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/contract"
)

func (s *UserUseCase) Detail(ctx context.Context, req userContract.DetailRequest) (res *userContract.DetailResponse, err error) {
	user, err := s.userRepo.GetByID(ctx, req.ID)
	if err != nil {
		return
	}
	userResponse := userContract.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	res = &userContract.DetailResponse{User: userResponse}
	return
}
