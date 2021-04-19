package rest

import (
	"rastebazaar/pkg/business"

	"github.com/go-chi/chi"
)

// MarketRouter : handle all market routes
func MarketRouter(router chi.Router, marketHandler *business.MarketHandler) chi.Router {

	router.Get("/api/v1/market", marketHandler.MarketerLogin)

	return router
}
