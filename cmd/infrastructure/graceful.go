package infrastructure

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (infra *Service) StopGracefully(server *http.Server, serverErr chan error) {
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-shutdownChannel:
		infra.Logger.Info("Caught signal ", sig, " Stop Gracefully")

		timeoutCfg := infra.Config.Graceful.TimeoutInSecond
		timeout := time.Duration(timeoutCfg) * time.Second

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		defer infra.CloseConnection()
		if err := server.Shutdown(ctx); err != nil {
			server.Close()
		}
	case err := <-serverErr:
		if err != nil {
			infra.Logger.Fatalf("server: %v", err)
		}
	}
}

func (infra *Service) CloseConnection() {
	infra.Logger.Info("closing connections...")

	if infra.DB != nil {
		infra.DB.Close()
	}
	if infra.Redis != nil {
		infra.Redis.Close()
	}
	if infra.Logger != nil {
		infra.Logger.CloseStream()
	}
}
