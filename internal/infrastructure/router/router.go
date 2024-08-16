package router

import (
	"fiber-app-paafff/internal/handlers"
	"fiber-app-paafff/internal/routes"

	"github.com/gofiber/fiber/v2"
)

// func SetupRoutes(app *fiber.App) {
//     routes.UserRoutes(app)
// }

func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	routes.UserRoutes(app, userHandler)
}
