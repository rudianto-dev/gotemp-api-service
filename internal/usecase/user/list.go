package user

import (
	"context"

	userContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/contract"
)

func (s *UserUseCase) List(ctx context.Context, req userContract.ListRequest) (res *userContract.ListResponse, err error) {

	users, err := s.userRepo.List(ctx)
	if err != nil {
		return
	}

	usersResponse := []userContract.UserResponse{}
	for _, user := range users {
		usersResponse = append(usersResponse, userContract.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	res = &userContract.ListResponse{Users: usersResponse}
	return
}
