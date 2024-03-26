package user

import userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"

// USER USE-CASE
type UserRequest struct {
	Name string `json:"name" validate:"required"`
}
type UserResponse struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	RoleType  userType.RoleType `json:"role_type" validate:"required"`
	UpdatedAt int64             `json:"updated_at"`
	CreatedAt int64             `json:"created_at"`
}

type DetailRequest struct {
	ID string `json:"id" validate:"required"`
}

type DetailResponse struct {
	User UserResponse `json:"user"`
}

type ListRequest struct {
}

type ListResponse struct {
	Users []UserResponse `json:"users"`
}

type CreateRequest struct {
	Name     string            `json:"name" validate:"required"`
	Username string            `json:"username" validate:"required"`
	RoleType userType.RoleType `json:"role_type" validate:"required"`
}

type CreateResponse struct {
	User UserResponse `json:"user"`
}

type UpdateRequest struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateResponse struct {
	User UserResponse `json:"user"`
}

type DeleteRequest struct {
	ID string `json:"id" validate:"required"`
}

type DeleteResponse struct {
	User UserResponse `json:"user"`
}
