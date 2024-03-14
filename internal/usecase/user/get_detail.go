package user

import (
	"context"

	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
)

func (s *UserUseCase) GetDetail(ctx context.Context, req userContract.GetDetailUserRequest) (*userContract.GetDetailUserResponse, error) {
	response := userContract.GetDetailUserResponse{}
	user, err := s.userRepo.GetByID(ctx, req.ID)
	if err != nil {
		return &response, err
	}
	response.User = userContract.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &response, nil
}
