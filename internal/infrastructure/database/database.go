package database

import (
	"fmt"
	"log"
	"os"

	"fiber-app-paafff/internal/domain/models" // Import the package that contains the "models" module

	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dbTimeZone := os.Getenv("DB_TIMEZONE")

	// Create DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode, dbTimeZone)

	// Connect to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// AutoMigrate will create the table based on the User struct
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	log.Println("Database migrated successfully")
}

// func InitDatabase() {
//     // Load .env file
//     err := godotenv.Load()
//     if err != nil {
//         log.Fatal("Error loading .env file")
//     }

// 		    // Get environment variables
// 				dbHost := os.Getenv("DB_HOST")
// 				dbUser := os.Getenv("DB_USER")
// 				dbPassword := os.Getenv("DB_PASSWORD")
// 				dbName := os.Getenv("DB_NAME")
// 				dbPort := os.Getenv("DB_PORT")
// 				dbSSLMode := os.Getenv("DB_SSLMODE")
// 				dbTimeZone := os.Getenv("DB_TIMEZONE")

// 	dsn := "host=localhost user=postgres password=paafff dbname=db_gofiber port=5432 sslmode=disable TimeZone=Asia/Shanghai"
// 	var err error
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("failed to connect database: ", err)
// 	}

// 	// AutoMigrate will create the table based on the User struct
// 	err = DB.AutoMigrate(&models.User{})
// 	if err != nil {
// 		log.Fatal("failed to migrate database: ", err)
// 	}

// 	log.Println("Database migrated successfully")
// }
