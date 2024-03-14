package util

import (
	"net/http"

	utilDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/util"
	res "github.com/rudianto-dev/gotemp-sdk/pkg/response"
)

func (s *UtilHandler) GetHealthStatus(w http.ResponseWriter, r *http.Request) {
	resp := utilDomain.GetHealthResponse{
		Status:  http.StatusOK,
		Message: "API Service is ready !",
	}
	res.Yay(w, r, http.StatusOK, resp)
}
