// package database

// import (
// 	"fmt"
// 	"log"

// 	"fiber-app-paafff/internal/config"
// 	"fiber-app-paafff/internal/domain/models"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func InitDatabase() {
// 	// Load database configuration from config.AppConfig
// 	dbConfig := config.AppConfig.Database

// 	// Create DSN
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
// 		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Name, dbConfig.Port, dbConfig.SSLMode, dbConfig.TimeZone)

// 	// Connect to the database
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

package database

import (
	"fmt"
	"log"
	"time"

	"fiber-app-paafff/internal/config"
	"fiber-app-paafff/internal/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	// Load database configuration from config.AppConfig
	dbConfig := config.AppConfig.Database

	fmt.Println(dbConfig)

	// Create DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Name, dbConfig.Port, dbConfig.SSLMode, dbConfig.TimeZone)

	var err error

	// Logika retry untuk koneksi database
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("failed to connect to database (attempt %d): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("could not connect to the database: ", err)
	}

	// AutoMigrate will create the table based on the User struct
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	log.Println("Database migrated successfully")
}
