package module

import (
	"github.com/go-chi/chi"
	utilHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/util"
)

func (m *Module) UtilRouting() *chi.Mux {
	handler := utilHandler.NewHandler(m.Infra.Logger)

	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		router.Post("/health", handler.GetHealthStatus)
	})
	return router
}
