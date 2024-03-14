package user

import (
	"context"

	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
)

func (s *UserUseCase) CreateUser(ctx context.Context, req userContract.CreateUserRequest) (*userContract.CreateUserResponse, error) {
	response := &userContract.CreateUserResponse{}
	user, err := userDomain.NewUser(userContract.CreateUser{
		Name: req.Name,
	})
	if err != nil {
		return response, err
	}
	user, err = s.userRepo.Create(ctx, nil, user)
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
