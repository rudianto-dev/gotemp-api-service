package user

import (
	"context"

	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
)

func (s *UserUseCase) DeleteUser(ctx context.Context, req userContract.DeleteUserRequest) (*userContract.DeleteUserResponse, error) {
	response := &userContract.DeleteUserResponse{}
	user, err := s.userRepo.GetByID(ctx, req.ID)
	if err != nil {
		return response, err
	}
	_, err = s.userRepo.Delete(ctx, nil, user)
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
