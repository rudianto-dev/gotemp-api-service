package user

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	userContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/contract"
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
)

type Repository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, user *userRepository.UserEntity) error
	Update(ctx context.Context, tx *sqlx.Tx, user *userRepository.UserEntity) error
	Delete(ctx context.Context, tx *sqlx.Tx, ID string) error
	GetByID(ctx context.Context, ID string) (*userDomain.User, error)
	List(ctx context.Context) ([]*userDomain.User, error)
}

type UseCase interface {
	Create(ctx context.Context, req userContract.CreateRequest) (*userContract.CreateResponse, error)
	Update(ctx context.Context, req userContract.UpdateRequest) (*userContract.UpdateResponse, error)
	Delete(ctx context.Context, req userContract.DeleteRequest) (*userContract.DeleteResponse, error)
	Detail(ctx context.Context, req userContract.DetailRequest) (*userContract.DetailResponse, error)
	List(ctx context.Context, req userContract.ListRequest) (*userContract.ListResponse, error)
	// RegisterUser()
}

type HandlerAPI interface {
	Detail(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	// OnBoarding()
}
