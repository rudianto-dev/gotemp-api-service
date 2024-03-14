package user

import (
	"net/http"

	"github.com/go-chi/chi"
	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
	res "github.com/rudianto-dev/gotemp-sdk/pkg/response"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	request := userContract.DeleteUserRequest{
		ID: chi.URLParam(r, "id"),
	}
	err := utils.ValidateStruct(request)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}

	resp, err := s.userUseCase.DeleteUser(ctx, request)
	if err != nil {
		s.logger.InfoWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}
	res.Yay(w, r, http.StatusOK, resp)
}
