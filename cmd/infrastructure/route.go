package infrastructure

import (
	"github.com/go-chi/chi"
	chim "github.com/go-chi/chi/middleware"
	"github.com/rudianto-dev/gotemp-api-service/internal/module"
	"github.com/rudianto-dev/gotemp-sdk/pkg/middleware"
)

func (srv *Service) CreateRouting() *chi.Mux {
	module := module.NewModule(&module.Service{
		Config: srv.Config,
		Logger: srv.Logger,
		DB:     srv.DB,
		Cache:  srv.Cache,
		JWT:    srv.JWT,
	})

	router := chi.NewRouter()
	router.Use(
		chim.NoCache,
		chim.RedirectSlashes,
		chim.RequestID,
		chim.Recoverer,
		chim.RealIP,
		chim.Heartbeat("/ping"),
		middleware.JWT(srv.JWT),
		// middleware.RequestLogger(srv.Config.NewLogrus()),
	)

	router.Route("/", func(router chi.Router) {
		router.Mount("/external", module.ExternalRoute())
		router.Mount("/internal", module.InternalRoute())
	})
	return router
}
