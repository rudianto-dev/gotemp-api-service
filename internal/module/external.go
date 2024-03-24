package module

import (
	"github.com/go-chi/chi"
	authHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/auth"
	userHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/user"
	utilHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/util"
	authUseCase "github.com/rudianto-dev/gotemp-api-service/internal/usecase/auth"
	userUseCase "github.com/rudianto-dev/gotemp-api-service/internal/usecase/user"
)

func (m *Module) ExternalRoute() *chi.Mux {
	utilHandler := utilHandler.NewHandler(m.Infra.Logger)

	authUseCase := authUseCase.NewUseCase(m.Infra.Logger, m.Infra.DB, m.UserRepository)
	authHandler := authHandler.NewHandler(m.Infra.Logger, authUseCase)

	userUseCase := userUseCase.NewUseCase(m.Infra.Logger, m.Infra.DB, m.UserRepository)
	userHandler := userHandler.NewHandler(m.Infra.Logger, userUseCase)

	router := chi.NewRouter()
	router.Route("/v1", func(router chi.Router) {
		router.Post("/health", utilHandler.GetHealthStatus)

		router.Route("/user", func(router chi.Router) {
			router.Post("/onboarding", userHandler.Onboarding)
		})
		router.Route("/auth", func(router chi.Router) {
			router.Get("/account/{username}", authHandler.CheckAccount)
		})
	})
	return router
}
