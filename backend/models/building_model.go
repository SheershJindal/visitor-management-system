package models

import "time"

type Building struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255);not null"`
	Address   string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Residents []Resident `gorm:"foreignKey:BuildingID"`
}
