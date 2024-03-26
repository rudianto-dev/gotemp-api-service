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

		router.Route("/client", func(router chi.Router) {
			router.Post("/list", m.ClientHandler.List)
			router.Get("/{id}", m.ClientHandler.Detail)
			router.Post("/", m.ClientHandler.Create)
			router.Put("/{id}", m.ClientHandler.Update)
			router.Delete("/{id}", m.ClientHandler.Delete)
		})
	})
	return router
}
