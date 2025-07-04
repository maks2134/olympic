package models

import "time"

type TestCase struct {
	ID       uint   `gorm:"primary_key"`
	TaskID   uint   `gorm:"not null"`
	Input    string `gorm:"type:text;not null"`
	Output   string `gorm:"type:text;not null"`
	IsHidden bool   `gorm:"default:false"`
}

type Solution struct {
	ID          uint      `gorm:"primary_key"`
	TaskID      uint      `gorm:"not null"`
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	Code        string    `gorm:"type:text;not null"`
	Language    string    `gorm:"not null"`
	Status      string    `gorm:"not null"`
	SubmittedAt time.Time `gorm:"not null"`
}
