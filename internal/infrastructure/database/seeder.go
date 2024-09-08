package database

import (
    "log"

    "github.com/bxcodec/faker/v3"
    "gorm.io/gorm"
    "fiber-app-paafff/internal/domain/models"
)

func SeedUsers(db *gorm.DB, count int) {
    for i := 0; i < count; i++ {
        user := models.User{
            Name:     faker.Name(),
            Email:    faker.Email(),
            Password: faker.Password(),
        }
        if err := db.Create(&user).Error; err != nil {
            log.Fatalf("Could not seed user: %v", err)
        }
    }
    log.Printf("Seeded %d users", count)
}