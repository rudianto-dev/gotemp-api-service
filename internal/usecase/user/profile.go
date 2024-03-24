package user

import (
	"context"

	userContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/contract"
)

func (s *UserUseCase) Profile(ctx context.Context, req userContract.ProfileRequest) (res *userContract.ProfileResponse, err error) {
	user, err := s.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return
	}
	phones, err := s.userRepo.ListPhoneNumberGetByUserID(ctx, req.UserID)
	if err != nil {
		return
	}
	phoneMap := []*userContract.Phone{}
	for _, p := range phones {
		phoneMap = append(phoneMap, &userContract.Phone{ID: p.ID, Number: p.PhoneNumber})
	}
	profile := userContract.Profile{
		ID:       user.ID,
		Name:     user.Name,
		RoleType: user.RoleType,
		Phones:   phoneMap,
	}
	res = &userContract.ProfileResponse{Profile: profile}
	return
}
