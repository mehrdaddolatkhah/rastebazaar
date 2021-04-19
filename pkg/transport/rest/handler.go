package rest

import (
	"context"
	"net/http"

	"rastebazaar/pkg/business"
	"rastebazaar/pkg/business/utils"

	"rastebazaar/pkg/infra"

	"rastebazaar/pkg/repository"

	routes "rastebazaar/pkg/transport/rest/routes"

	web "rastebazaar/pkg/transport/rest/routes/web"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

// RouteHandler handle all routes and accessibility here
func RouteHandler(infra *infra.Infrastructure) http.Handler {

	// handle cors for go-chi
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		ExposedHeaders: []string{"Content-Type", "Set-Cookie", "Cookie"},
		//AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "application/x-www-form-urlencoded"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	//tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
	// define a new admin handler and repository
	adminRepo := repository.NewAdminRepo(infra)
	adminHandler := business.NewAdminHandler(adminRepo, infra.Template)

	categoryRepo := repository.NewCategoryRepo(infra)
	categoryHandler := business.NewCategoryHandler(categoryRepo, infra.Template)

	// define a new user handler and repository
	userRepo := repository.NewUserRepo(infra)
	userHandler := business.NewUserHandler(userRepo)

	// define a new market handler and repository
	marketRepo := repository.NewMarketRepo(infra)
	marketHandler := business.NewMarketHandler(marketRepo)

	// define a new market handler and repository
	authRepo := repository.NewAuthRepo(infra)
	authHandler := business.NewAuthHandler(authRepo)

	router := chi.NewRouter()

	router.Use(cors.Handler)
	router.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	// Protected routes
	router.Group(func(r chi.Router) {

		// check user role and access here
		r.Use(rasteBazaarMiddleware)
		//r.Use(jwtauth.Verifier(tokenAuth))
		//r.Use(jwtauth.Authenticator)

		// Mount the admin sub-router
		r.Mount("/secure/admin", routes.AdminRouter(router, adminHandler))

		r.Mount("/secure/category", web.CategoryWebRouter(router, categoryHandler, infra.Template))

		// Mount the market sub-router
		r.Mount("/api/v1/market", routes.MarketRouter(router, marketHandler))

		// Mount the user sub-router
		r.Mount("/api/v1/user", routes.UserRouter(router, userHandler))

	})

	// Public routes
	router.Group(func(r chi.Router) {

		r.Mount("/api/v1", routes.PublicRouter(router, authHandler, adminHandler))

		r.Mount("/admin", web.AdminWebRouter(router, adminHandler, infra.Template))
	})

	return router
}

func rasteBazaarMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		accessToken, err := r.Cookie("token")

		if err != nil {
			panic(err)
		}

		if accessToken.Value == "" {
			return
		}

		userID := utils.TokenExtractor(accessToken.Value)

		if err != nil {
			panic(err)
		}

		ctx := context.WithValue(r.Context(), "userId", userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
