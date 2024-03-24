package user

import (
	"net/http"

	userContract "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/contract"
	"github.com/rudianto-dev/gotemp-sdk/pkg/middleware"
	res "github.com/rudianto-dev/gotemp-sdk/pkg/response"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserHandler) Profile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	claim := middleware.GetClaims(r)
	request := userContract.ProfileRequest{
		UserID: claim.ID,
	}
	err := utils.ValidateStruct(request)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}

	resp, err := s.userUseCase.Profile(ctx, request)
	if err != nil {
		s.logger.InfoWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}
	res.Yay(w, r, http.StatusOK, resp)
}
