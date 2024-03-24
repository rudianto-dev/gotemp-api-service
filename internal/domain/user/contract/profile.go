package user

import (
	userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"
)

type Profile struct {
	ID       string            `json:"id" validate:"required"`
	Name     string            `json:"name" validate:"required"`
	Phones   []*Phone          `json:"phones" validate:"required"`
	RoleType userType.RoleType `json:"role_type" validate:"required"`
}

type Phone struct {
	ID     string `json:"id" validate:"required"`
	Number string `json:"number" validate:"required"`
}

type ProfileRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

type ProfileResponse struct {
	Profile Profile `json:"profile"`
}
