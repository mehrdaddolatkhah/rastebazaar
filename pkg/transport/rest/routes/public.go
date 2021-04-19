package rest

import (
	"rastebazaar/pkg/business"

	"github.com/go-chi/chi"
)

// PublicRouter : handle all user routes
func PublicRouter(router chi.Router, userHandler *business.AuthHandler, adminHandler *business.AdminHandler) chi.Router {

	// Register and Login for admin
	router.Post("/register/admin", adminHandler.Register)
	router.Post("/login/admin", adminHandler.Login)

	// Login and Verify for marketer
	router.Post("/login/market", userHandler.Verify)
	router.Post("/verify/market", userHandler.Verify)

	// Login and Verify for normal users
	router.Post("/login", userHandler.Login)
	router.Post("/verify", userHandler.Verify)

	return router
}
