package infrastructure

import (
	"net/http"

	"github.com/go-chi/chi"
	chim "github.com/go-chi/chi/middleware"
	"github.com/rudianto-dev/gotemp-api-service/internal/module"
	"github.com/rudianto-dev/gotemp-sdk/pkg/middleware"
)

func (srv *Service) RunAPI() {
	router := srv.SetupAPI()
	
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		srv.Logger.Infof("%s %s", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		srv.Logger.Fatal(err)
	}

	server := &http.Server{Addr: srv.Config.Host.Address, Handler: router}
	srv.Logger.Infof("API GW serving at %s", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			srv.Logger.Error(err)
		}
	}()
	srv.StopGracefully(server)
}

func (srv *Service) SetupAPI() *chi.Mux {
	module := module.NewModule(&module.Service{
		Config: srv.Config,
		Logger: srv.Logger,
		DB:     srv.DB,
		Redis:  srv.Redis,
	})
	utilHandlerAPI := module.UtilHandlerAPI()
	userHandlerAPI := module.UserHandlerAPI()

	router := chi.NewRouter()
	router.Use(
		chim.NoCache,
		chim.RedirectSlashes,
		chim.RequestID,
		chim.Recoverer,
		chim.RealIP,
		chim.Heartbeat("/ping"),
		middleware.RequestLogger(srv.Config.NewLogrus()),
	)
	router.Route("/", func(r chi.Router) {
		router.Get("/health", utilHandlerAPI.GetHealthStatus)
		router.Route("/v1", func(r chi.Router) {
			router.Route("/user", func(r chi.Router) {
				router.Post("/list", userHandlerAPI.List)
				router.Get("/{id}", userHandlerAPI.GetDetail)
				router.Post("/", userHandlerAPI.Create)
				router.Put("/{id}", userHandlerAPI.Update)
				router.Delete("/{id}", userHandlerAPI.Delete)
			})
		})
	})
	return router
}
