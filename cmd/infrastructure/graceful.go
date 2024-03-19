package infrastructure

import (
	"context"
	"net/http"
	"time"
	"github.com/rudianto-dev/gotemp-sdk/pkg/gracefully"
)

func (srv *Service) StopGracefully(server *http.Server) {
	gracefully.WaitAndShutdown(map[string]any{
		"database": func() error {
			if err := srv.DB.Close(); err != nil {
				return err
			}
			return nil
		},
		"redis": func() error {
			if err := srv.Redis.Close(); err != nil {
				return err
			}
			return nil
		},
		"server": func() {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			server.Shutdown(ctx)
		},
	})
}
