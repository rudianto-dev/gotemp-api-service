package util

import (
	"net/http"
)

type IUseCase interface {
}

type IHandlerAPI interface {
	GetHealthStatus(w http.ResponseWriter, r *http.Request)
}
