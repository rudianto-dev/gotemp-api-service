package user

import (
	"context"

	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
)

func (s *UserUseCase) GetProfile(ctx context.Context, req userContract.GetProfileRequest) (*userContract.GetProfileResponse, error) {
	userProfile := userContract.GetProfileResponse{}
	user, err := s.userRepo.GetByID(ctx, req.ID)
	if err != nil {
		return &userProfile, err
	}
	userProfile.User = userContract.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &userProfile, nil
}
