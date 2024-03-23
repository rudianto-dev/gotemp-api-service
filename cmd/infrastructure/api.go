package infrastructure

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (srv *Service) RunAPI() {
	router := srv.CreateRouting()

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
