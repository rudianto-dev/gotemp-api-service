package user

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	userContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/contract"
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
	userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"
	userRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/repository"
)

type Repository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, userEntity *userRepository.UserEntity) error
	Update(ctx context.Context, tx *sqlx.Tx, userEntity *userRepository.UserEntity) error
	UpdatePassword(ctx context.Context, tx *sqlx.Tx, ID, password string) error
	UpdateStatus(ctx context.Context, tx *sqlx.Tx, ID string, status userType.Status) error
	Delete(ctx context.Context, tx *sqlx.Tx, ID string) error
	GetByID(ctx context.Context, ID string) (*userDomain.User, error)
	GetByUsername(ctx context.Context, username string) (*userDomain.User, error)
	List(ctx context.Context) ([]*userDomain.User, error)

	GetCIFByReferenceID(ctx context.Context, referenceID string) (*userDomain.CIF, error)
	InsertCIF(ctx context.Context, tx *sqlx.Tx, CIFEntity *userRepository.CIFEntity) error
	InsertPhoneNumber(ctx context.Context, tx *sqlx.Tx, PhoneNumberEntity *userRepository.PhoneNumberEntity) error
	InsertEmail(ctx context.Context, tx *sqlx.Tx, EmailEntity *userRepository.EmailEntity) error
	InsertDeviceInstance(ctx context.Context, tx *sqlx.Tx, DeviceInstanceEntity *userRepository.DeviceInstanceEntity) error

	ListPhoneNumberGetByUserID(ctx context.Context, UserID string) ([]*userDomain.PhoneNumber, error)
}

type UseCase interface {
	Create(ctx context.Context, req userContract.CreateRequest) (*userContract.CreateResponse, error)
	Update(ctx context.Context, req userContract.UpdateRequest) (*userContract.UpdateResponse, error)
	Delete(ctx context.Context, req userContract.DeleteRequest) (*userContract.DeleteResponse, error)
	Detail(ctx context.Context, req userContract.DetailRequest) (*userContract.DetailResponse, error)
	List(ctx context.Context, req userContract.ListRequest) (*userContract.ListResponse, error)

	Onboarding(ctx context.Context, req userContract.OnboardingRequest) (*userContract.OnboardingResponse, error)
	Profile(ctx context.Context, req userContract.ProfileRequest) (*userContract.ProfileResponse, error)
}

type HandlerAPI interface {
	Detail(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)

	Onboarding(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
}
