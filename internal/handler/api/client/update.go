package client

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	clientContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/client/contract"
	res "github.com/rudianto-dev/gotemp-sdk/pkg/response"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *ClientHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	request := clientContract.UpdateRequest{}
	if err := render.Decode(r, &request); err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}
	request.ID = chi.URLParam(r, "id")
	err := utils.ValidateStruct(request)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}

	resp, err := s.clientUseCase.Update(ctx, request)
	if err != nil {
		s.logger.InfoWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}
	res.Yay(w, r, http.StatusOK, resp)
}
