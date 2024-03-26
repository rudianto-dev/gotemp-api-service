package client

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	clientContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/contract"
	clientDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/model"
	clientRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/repository"
)

type Repository interface {
	Insert(ctx context.Context, tx *sqlx.Tx, clientEntity *clientRepository.ClientEntity) error
	Update(ctx context.Context, tx *sqlx.Tx, clientEntity *clientRepository.ClientEntity) error
	Delete(ctx context.Context, tx *sqlx.Tx, id string) error
	GetByID(ctx context.Context, id string) (*clientDomain.Client, error)
	GetByClientID(ctx context.Context, clientID string) (*clientDomain.Client, error)
	List(ctx context.Context) ([]*clientDomain.Client, error)

	SaveCache(ctx context.Context, req *clientRepository.CacheEntity) error
	GetCache(ctx context.Context, clientID string) (string, error)
	DeleteCache(ctx context.Context, clientID string) error
}

type UseCase interface {
	Create(ctx context.Context, req clientContract.CreateRequest) (*clientContract.CreateResponse, error)
	Update(ctx context.Context, req clientContract.UpdateRequest) (*clientContract.UpdateResponse, error)
	Delete(ctx context.Context, req clientContract.DeleteRequest) (*clientContract.DeleteResponse, error)
	Detail(ctx context.Context, req clientContract.DetailRequest) (*clientContract.DetailResponse, error)
	List(ctx context.Context, req clientContract.ListRequest) (*clientContract.ListResponse, error)
}

type HandlerAPI interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Detail(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}
