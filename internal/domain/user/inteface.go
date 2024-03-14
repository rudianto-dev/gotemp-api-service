package user

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
)

type IRepository interface {
	GetByID(ctx context.Context, id string) (*User, error)
	GetList(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, tx *sqlx.Tx, user *User) (*User, error)
	Update(ctx context.Context, tx *sqlx.Tx, user *User) (*User, error)
	Delete(ctx context.Context, tx *sqlx.Tx, user *User) (*User, error)
}

type IUseCase interface {
	GetProfile(ctx context.Context, req userContract.GetProfileRequest) (*userContract.GetProfileResponse, error)
	GetList(ctx context.Context, req userContract.GetListRequest) (*userContract.GetListResponse, error)
	CreateUser(ctx context.Context, req userContract.CreateUserRequest) (*userContract.CreateUserResponse, error)
	UpdateUser(ctx context.Context, req userContract.UpdateUserRequest) (*userContract.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req userContract.DeleteUserRequest) (*userContract.DeleteUserResponse, error)
}

type IHandlerAPI interface {
	GetProfile(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
