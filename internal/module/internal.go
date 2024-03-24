package module

import (
	"github.com/go-chi/chi"
	userHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/user"
	userUseCase "github.com/rudianto-dev/gotemp-api-service/internal/usecase/user"
)

func (m *Module) InternalRoute() *chi.Mux {
	useCase := userUseCase.NewUseCase(m.Infra.Logger, m.Infra.DB, m.UserRepository)
	handler := userHandler.NewHandler(m.Infra.Logger, useCase)

	router := chi.NewRouter()
	router.Route("/v1", func(router chi.Router) {
		router.Route("/user", func(router chi.Router) {
			router.Post("/list", handler.List)
			router.Get("/{id}", handler.Detail)
			router.Post("/", handler.Create)
			router.Put("/{id}", handler.Update)
			router.Delete("/{id}", handler.Delete)
		})
	})
	return router
}
