package auth

import (
	"net/http"

	"github.com/go-chi/render"
	authContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/contract"
	res "github.com/rudianto-dev/gotemp-sdk/pkg/response"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	request := authContract.ResetPasswordRequest{}
	if err := render.Decode(r, &request); err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}
	err := utils.ValidateStruct(request)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}

	resp, err := s.authUseCase.ResetPassword(ctx, request)
	if err != nil {
		s.logger.InfoWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}
	res.Yay(w, r, http.StatusOK, resp)
}
