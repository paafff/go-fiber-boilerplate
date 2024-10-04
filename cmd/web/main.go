package main

import (
	"fiber-app-paafff/internal/auth"
	"fiber-app-paafff/internal/config"
	"fiber-app-paafff/internal/domain/services"
	"fiber-app-paafff/internal/handlers"
	"fiber-app-paafff/internal/infrastructure/database"
	"fiber-app-paafff/internal/infrastructure/router"
	"fiber-app-paafff/internal/repositories"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	config.LoadConfig()
	log.Println("Configuration loaded")

	// Inisialisasi database
	database.InitDatabase()
	log.Println("Database initialized")

	// Handle command line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "seed":
			database.SeedUsers(database.DB, 10)
			fmt.Println("Seeded 10 users")
			return
		case "reset-db":
			database.ResetDatabase(database.DB)
			fmt.Println("Database reset.")
			return
		}
	}

	// Inisialisasi repository
	userRepository := repositories.NewUserRepository(database.DB)
	log.Println("User repository initialized")

	// Inisialisasi service
	userService := services.NewUserService(userRepository)
	authService := auth.NewAuthService(userRepository)
	log.Println("User service and Auth service initialized")

	// Inisialisasi handler
	userHandler := handlers.NewUserHandler(userService)
	authHandler := auth.NewAuthHandler(authService)
	log.Println("User handler and Auth handler initialized")

	// // Inisialisasi service
	// userService := services.NewUserService(userRepository)
	// log.Println("User service initialized")

	// // Inisialisasi handler
	// userHandler := handlers.NewUserHandler(userService)
	// log.Println("User handler initialized")

	// Inisialisasi Fiber
	app := fiber.New()
	log.Println("Fiber app initialized")

	// Setup routes
	router.SetupRoutes(app, userHandler, authHandler)
	log.Println("Routes set up")

	// Jalankan server
	log.Println("Server running on port 3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
