package module

import (
	"github.com/go-chi/chi"
	authHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/auth"
	otpHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/otp"
	userHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/user"
	utilHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/util"
	authUseCase "github.com/rudianto-dev/gotemp-api-service/internal/usecase/auth"
	otpUseCase "github.com/rudianto-dev/gotemp-api-service/internal/usecase/otp"
	userUseCase "github.com/rudianto-dev/gotemp-api-service/internal/usecase/user"
)

func (m *Module) ExternalRoute() *chi.Mux {
	authUseCase := authUseCase.NewUseCase(m.Infra.Logger, m.Infra.DB, m.UserRepository)
	userUseCase := userUseCase.NewUseCase(m.Infra.Logger, m.Infra.DB, m.UserRepository)
	otpUseCase := otpUseCase.NewUseCase(m.Infra.Logger, m.OTPRepository)

	utilHandler := utilHandler.NewHandler(m.Infra.Logger)
	authHandler := authHandler.NewHandler(m.Infra.Logger, authUseCase)
	userHandler := userHandler.NewHandler(m.Infra.Logger, userUseCase)
	otpHandler := otpHandler.NewHandler(m.Infra.Logger, otpUseCase)

	router := chi.NewRouter()
	router.Route("/v1", func(router chi.Router) {
		router.Post("/health", utilHandler.GetHealthStatus)

		router.Route("/user", func(router chi.Router) {
			router.Post("/onboarding", userHandler.Onboarding)
		})
		router.Route("/auth", func(router chi.Router) {
			router.Get("/account/{username}", authHandler.CheckAccount)
			router.Post("/login", authHandler.CheckAccount)
			router.Post("/logout", authHandler.CheckAccount)
			router.Post("/refresh-token", authHandler.CheckAccount)
		})
		router.Route("/password", func(router chi.Router) {
			router.Post("/check", authHandler.CheckAccount)
			router.Post("/reset", authHandler.CheckAccount)
		})
		router.Route("/otp", func(router chi.Router) {
			router.Post("/send", otpHandler.SendOTP)
			router.Post("/verify", otpHandler.VerifyOTP)
		})
	})
	return router
}
