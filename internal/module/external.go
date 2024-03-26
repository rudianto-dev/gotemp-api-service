package module

import (
	"github.com/go-chi/chi"
	"github.com/rudianto-dev/gotemp-sdk/pkg/middleware"
)

func (m *Module) ExternalRoute() *chi.Mux {
	router := chi.NewRouter()
	router.Route("/v1", func(router chi.Router) {
		// util section
		router.Post("/health", m.UtilHandler.GetHealthStatus)
		router.Post("/build", m.UtilHandler.BuildRequest)
		// user section
		router.Route("/user", func(router chi.Router) {
			router.Group(func(router chi.Router) {
				router.Use(middleware.ClientAuthenticated())
				router.Post("/onboarding", m.UserHandler.Onboarding)
			})
			router.Group(func(router chi.Router) {
				router.Use(middleware.GuardAuthenticated(middleware.TokenFromHeader))
				router.Get("/profile", m.UserHandler.Profile)
			})
		})
		// auth section
		router.Route("/auth", func(router chi.Router) {
			router.Get("/account/{username}", m.AuthHandler.CheckAccount)
			router.Post("/refresh-token", m.AuthHandler.RefreshToken)
			router.Post("/login", m.AuthHandler.Login)
			router.Group(func(router chi.Router) {
				router.Use(middleware.GuardAuthenticated(middleware.TokenFromHeader))
				router.Post("/logout", m.AuthHandler.Logout)
			})
		})
		// password section
		router.Route("/password", func(router chi.Router) {
			router.Post("/reset", m.AuthHandler.ResetPassword)
			router.Group(func(router chi.Router) {
				router.Use(middleware.GuardAuthenticated(middleware.TokenFromHeader))
				router.Post("/check", m.AuthHandler.CheckPassword)
			})
		})
		// otp section
		router.Route("/otp", func(router chi.Router) {
			router.Post("/send", m.OTPHandler.SendOTP)
			router.Post("/verify", m.OTPHandler.VerifyOTP)
		})
	})
	return router
}
