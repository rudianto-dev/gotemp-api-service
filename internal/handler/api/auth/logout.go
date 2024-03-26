package auth

import (
	"net/http"

	authContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/auth/contract"
	"github.com/rudianto-dev/gotemp-sdk/pkg/middleware"
	res "github.com/rudianto-dev/gotemp-sdk/pkg/response"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	claim := middleware.GetClaims(r)
	request := authContract.LogoutRequest{
		TokenID: claim.ID,
	}
	err := utils.ValidateStruct(request)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}

	resp, err := s.authUseCase.Logout(ctx, request)
	if err != nil {
		s.logger.InfoWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}
	res.Yay(w, r, http.StatusOK, resp)
}
