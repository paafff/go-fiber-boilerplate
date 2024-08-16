package models

import (
	"time"
)

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255"`
	Description string `gorm:"size:255"`
	Price       float64
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
