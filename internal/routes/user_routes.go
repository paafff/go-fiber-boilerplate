package routes

import (
    "github.com/gofiber/fiber/v2"
    "fiber-app-paafff/internal/handlers"
)

func UserRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
  app.Get("/user/:id", userHandler.GetUser)
}

