package module

import (
	"github.com/go-chi/chi"
)

func (m *Module) InternalRoute() *chi.Mux {
	router := chi.NewRouter()
	router.Route("/v1", func(router chi.Router) {
		router.Route("/user", func(router chi.Router) {
			router.Post("/list", m.UserHandler.List)
			router.Get("/{id}", m.UserHandler.Detail)
			router.Post("/", m.UserHandler.Create)
			router.Put("/{id}", m.UserHandler.Update)
			router.Delete("/{id}", m.UserHandler.Delete)
		})
	})
	return router
}
