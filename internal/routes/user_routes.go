package routes

import (
	"fiber-app-paafff/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	app.Get("/user/:id", userHandler.GetUser)
	app.Post("/user", userHandler.CreateUser)
	app.Put("/user/:id", userHandler.UpdateUser)
	app.Delete("/user/:id", userHandler.DeleteUser)
	app.Get("/users", userHandler.ListUsers)
}
