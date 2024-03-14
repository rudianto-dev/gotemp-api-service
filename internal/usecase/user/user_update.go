package user

import (
	"context"

	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
)

func (s *UserUseCase) UpdateUser(ctx context.Context, req userContract.UpdateUserRequest) (*userContract.UpdateUserResponse, error) {
	response := &userContract.UpdateUserResponse{}
	user, err := s.userRepo.GetByID(ctx, req.ID)
	if err != nil {
		return response, err
	}
	user.Name = req.Name
	user, err = s.userRepo.Update(ctx, nil, user)
	if err != nil {
		return response, err
	}
	response.User = userContract.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}
