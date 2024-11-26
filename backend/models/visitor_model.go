package models

import "time"

type Visitor struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"type:varchar(255);not null"`
	VisitTime  time.Time
	ResidentID uint `gorm:"not null"`
}
