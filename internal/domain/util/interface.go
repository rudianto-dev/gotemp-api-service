package util

import (
	"net/http"
)

type IUseCase interface {
}

type HandlerAPI interface {
	GetHealthStatus(w http.ResponseWriter, r *http.Request)
}
