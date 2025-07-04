package models

import "time"

type Board struct {
	ID         uint     `gorm:"primary_key"`
	OlympiadID uint     `gorm:"not null"`
	Olympiad   Olympiad `gorm:"foreignKey:OlympiadID"`
	UserID     uint     `gorm:"not null"`
	User       User     `gorm:"foreignKey:UserID"`
	Score      uint     `gorm:"default:0"`
	Position   uint
	Solved     uint `gorm:"default:0"`
	LastSubmit time.Time
}
