package rest

import (
	"rastebazaar/pkg/business"

	"github.com/go-chi/chi"
)

// UserRouter : handle all user routes
func UserRouter(router chi.Router, userHandler *business.UserHandler) chi.Router {

	//router.Get("/api/v1/user", userHandler.HelloUser)

	return router
}
