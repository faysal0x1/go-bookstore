package routes

import (
	"net/http"

	"github.com/faysal0x1/go-bookstore/pkg/controllers"
	"github.com/faysal0x1/go-bookstore/pkg/middleware"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

type RouteParams struct {
	fx.In
	Router         *mux.Router
	AuthMiddleware *middleware.AuthMiddleware
	BookController *controllers.BookController
	AuthController *controllers.AuthController
}

func RegisterRoutes(p RouteParams) {
	// 1. Setup Middleware Registry (Laravel-like named middlewares)
	registry := map[string]interface{}{
		"auth": p.AuthMiddleware.Handler,
		"role": func(role string) MiddlewareFunc {
			return MiddlewareFunc(middleware.RoleMiddleware(role))
		},
	}

	// 2. Global Middlewares
	p.Router.Use(middleware.RateLimitMiddleware())

	// 3. Static Files (Laravel-like storage)
	p.Router.PathPrefix("/storage/").Handler(http.StripPrefix("/storage/", http.FileServer(http.Dir("./storage"))))

	// 4. Initialize the Fluent Router
	api := NewRouteGroup(p.Router, registry)

	// 5. Route Definitions (Like Laravel's api.php)

	// Public Routes
	api.Post("/register", p.AuthController.Register)
	api.Post("/login", p.AuthController.Login)

	// Protected Book Routes
	api.Middleware("auth").Prefix("/book").Group(func(book *RouteGroup) {
		book.Post("", p.BookController.CreateBook)
		book.Get("", p.BookController.GetBooks)
		book.Get("/{bookId}", p.BookController.GetBookByID)

		// Admin Only Nested Group
		book.Middleware("role:Admin").Prefix("/{bookId}").Group(func(admin *RouteGroup) {
			admin.Put("", p.BookController.UpdateBook)
			admin.Delete("", p.BookController.DeleteBook)
		})
	})
}
