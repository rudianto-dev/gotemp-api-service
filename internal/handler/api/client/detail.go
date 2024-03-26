package client

import (
	"net/http"

	"github.com/go-chi/chi"
	clientContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/contract"
	res "github.com/rudianto-dev/gotemp-sdk/pkg/response"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *ClientHandler) Detail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	request := clientContract.DetailRequest{
		ID: chi.URLParam(r, "id"),
	}
	err := utils.ValidateStruct(request)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}

	resp, err := s.clientUseCase.Detail(ctx, request)
	if err != nil {
		s.logger.InfoWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}
	res.Yay(w, r, http.StatusOK, resp)
}
