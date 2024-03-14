package infrastructure

import (
	"net/http"

	"github.com/go-chi/chi"
	chim "github.com/go-chi/chi/middleware"
	"github.com/rudianto-dev/gotemp-api-service/internal/module"
	"github.com/rudianto-dev/gotemp-sdk/pkg/middleware"
)

func (infra *Service) CreateAPIService() error {
	r := chi.NewRouter()
	r.Use(
		chim.NoCache,
		chim.RedirectSlashes,
		chim.RequestID,
		chim.Recoverer,
		chim.RealIP,
		chim.Heartbeat("/ping"),
		middleware.RequestLogger(infra.Config.NewLogrus()),
	)

	// load module
	module := module.NewModule(&module.Service{
		Config: infra.Config,
		Logger: infra.Logger,
		DB:     infra.DB,
		Redis:  infra.Redis,
	})
	userHandlerAPI := module.UserHandlerAPI()

	r.Route("/v1", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong !"))
		})
		r.Route("/user", func(r chi.Router) {
			r.Post("/", userHandlerAPI.Create)
			r.Put("/{id}", userHandlerAPI.Update)
			r.Get("/{id}", userHandlerAPI.GetProfile)
			r.Post("/list", userHandlerAPI.List)
			r.Delete("/{id}", userHandlerAPI.Delete)
		})
	})

	server := http.Server{
		Addr:    infra.Config.Host.Address,
		Handler: r,
	}
	serverErr := make(chan error, 1)
	go func() {
		infra.Logger.Infof("User API serving at %s", server.Addr)
		serverErr <- server.ListenAndServe()
	}()

	infra.StopGracefully(&server, serverErr)
	return nil
}
