package web

import (
	"html/template"

	"rastebazaar/pkg/business"

	"github.com/go-chi/chi"
)

// AdminWebRouter : handle all user routes
func AdminWebRouter(router chi.Router, adminHandler *business.AdminHandler, tmpl *template.Template) chi.Router {

	// Register and Login for admin
	router.Get("/register", adminHandler.GetAdminRegisterView)
	router.Get("/login", adminHandler.GetAdminLoginView)

	router.Post("/login", adminHandler.PostAdminLoginView)
	router.Post("/register", adminHandler.PostAdminRegisterView)

	return router
}
