package infrastructure

import (
	"github.com/go-chi/chi"
	chim "github.com/go-chi/chi/middleware"
	"github.com/rudianto-dev/gotemp-api-service/internal/module"
)

func (srv *Service) CreateRouting() *chi.Mux {
	module := module.NewModule(&module.Service{
		Config: srv.Config,
		Logger: srv.Logger,
		DB:     srv.DB,
		Redis:  srv.Redis,
	})

	router := chi.NewRouter()
	router.Use(
		chim.NoCache,
		chim.RedirectSlashes,
		chim.RequestID,
		chim.Recoverer,
		chim.RealIP,
		chim.Heartbeat("/ping"),
		// middleware.RequestLogger(srv.Config.NewLogrus()),
	)

	router.Route("/", func(router chi.Router) {
		router.Mount("/utility", module.UtilRouting())

		router.Route("/v1", func(router chi.Router) {
			router.Mount("/user", module.UserRouting())
		})
	})
	return router
}