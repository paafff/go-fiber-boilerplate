// package router

// import (
// 	"fiber-app-paafff/internal/handlers"
// 	"fiber-app-paafff/internal/routes"

// 	"github.com/gofiber/fiber/v2"
// )

// // func SetupRoutes(app *fiber.App) {
// //     routes.UserRoutes(app)
// // }

// func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
// 	routes.UserRoutes(app, userHandler)
// }

package router

import (
	"fiber-app-paafff/internal/auth"
	"fiber-app-paafff/internal/handlers"
	"fiber-app-paafff/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler, authHandler *auth.AuthHandler) {
	api := app.Group("/api")

	// Public routes
	api.Post("/login", authHandler.Login)
	api.Post("/register", authHandler.Register)

	// Protected routes
	api.Use(auth.AuthMiddleware)
	routes.UserRoutes(api, userHandler)
	api.Get("/users", userHandler.GetUsers)
	// Add other protected routes here
}
