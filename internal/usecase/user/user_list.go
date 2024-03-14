package user

import (
	"context"

	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
)

func (s *UserUseCase) GetList(ctx context.Context, req userContract.GetListUserRequest) (*userContract.GetListUserResponse, error) {
	response := userContract.GetListUserResponse{}
	users, err := s.userRepo.GetList(ctx)
	if err != nil {
		return &response, err
	}

	for _, user := range users {
		response.Users = append(response.Users, userContract.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return &response, nil
}
