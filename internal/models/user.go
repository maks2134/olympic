package models

import "time"

type User struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"default:'participant'"`
	Telegram  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
