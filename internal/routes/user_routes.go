package routes

import (
	"fiber-app-paafff/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, userHandler *handlers.UserHandler) {
	router.Get("/user/:id", userHandler.GetUser)
	router.Post("/user", userHandler.CreateUser)
	router.Put("/user/:id", userHandler.UpdateUser)
	router.Delete("/user/:id", userHandler.DeleteUser)
	// app.Get("/users", userHandler.ListUsers)
}
