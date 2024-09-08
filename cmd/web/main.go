package main

import (
	"fiber-app-paafff/internal/config"
	"fiber-app-paafff/internal/domain/services"
	"fiber-app-paafff/internal/handlers"
	"fiber-app-paafff/internal/infrastructure/database"
	"fiber-app-paafff/internal/infrastructure/router"
	"fiber-app-paafff/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Load configuration
	config.LoadConfig()

	// Inisialisasi database
	database.InitDatabase()

	// Seed dummy data
	database.SeedUsers(database.DB, 10) // Seed 10 users

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
