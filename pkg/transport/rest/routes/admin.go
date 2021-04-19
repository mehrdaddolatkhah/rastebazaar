package rest

import (
	"net/http"

	"rastebazaar/pkg/business"

	"github.com/go-chi/chi"
)

// AdminRouter todo : change it later. useless
// AdminRouter : handle all admin routes
func AdminRouter(router chi.Router, adminHandler *business.AdminHandler) http.Handler {
	router.Get("/hello", adminHandler.Register)

	router.Get("/panel", adminHandler.GetAdminPanelView)

	return router
}
