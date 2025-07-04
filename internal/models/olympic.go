package models

import "time"

type Olympiad struct {
	ID           uint      `gorm:"primary_key"`
	Title        string    `gorm:"not null"`
	Description  string    `gorm:"type:text"`
	StartTime    time.Time `gorm:"not null"`
	EndTime      time.Time `gorm:"not null"`
	IsActive     bool      `gorm:"default:false"`
	Tasks        []Task
	Boards       []Board
	Participants []User `gorm:"many2many:olympiad_participants;"`
}
