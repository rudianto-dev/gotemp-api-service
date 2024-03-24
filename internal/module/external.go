package module

import (
	"github.com/go-chi/chi"
	userHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/user"
	utilHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/util"
	userUseCase "github.com/rudianto-dev/gotemp-api-service/internal/usecase/user"
)

func (m *Module) ExternalRoute() *chi.Mux {
	utilHandler := utilHandler.NewHandler(m.Infra.Logger)

	userUseCase := userUseCase.NewUseCase(m.Infra.Logger, m.Infra.DB, m.UserRepository)
	userHandler := userHandler.NewHandler(m.Infra.Logger, userUseCase)

	router := chi.NewRouter()
	router.Route("/v1", func(router chi.Router) {
		router.Post("/health", utilHandler.GetHealthStatus)

		router.Route("/user", func(router chi.Router) {
			router.Post("/onboarding", userHandler.Onboarding)
		})
	})
	return router
}
