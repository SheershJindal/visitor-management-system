package models

import "time"

type Resident struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"type:varchar(255);not null"`
	Email      string `gorm:"type:varchar(255);unique;not null"`
	FlatNumber string `gorm:"type:varchar(50);not null"`
	BuildingID uint   `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
