package web

import (
	"html/template"
	"rastebazaar/pkg/business"

	"github.com/go-chi/chi"
)

// CategoryWebRouter : handle all user routes
func CategoryWebRouter(router chi.Router, categoryHandler *business.CategoryHandler, tmpl *template.Template) chi.Router {

	// Register and Login for admin
	router.Get("/category-list", categoryHandler.GetCategoryListView)

	router.Get("/add-category", categoryHandler.GetAddCategoryView)

	router.Post("/add-category", categoryHandler.PostAddCategoryView)

	return router
}
