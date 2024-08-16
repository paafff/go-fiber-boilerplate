package main

import (
    "fiber-app-paafff/internal/infrastructure/database"
    "fiber-app-paafff/internal/infrastructure/router"
    "fiber-app-paafff/internal/repositories"
    "fiber-app-paafff/internal/domain/services"
    "fiber-app-paafff/internal/handlers"
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Inisialisasi database
    database.InitDatabase()

    // Inisialisasi repository
    userRepository := repositories.NewUserRepository(database.DB)

    // Inisialisasi service
    userService := services.NewUserService(userRepository)

    // Inisialisasi handler
    userHandler := handlers.NewUserHandler(userService)

    // Inisialisasi Fiber
    app := fiber.New()

    // Setup routes
    router.SetupRoutes(app, userHandler)

    // Jalankan server
    app.Listen(":3000")
}