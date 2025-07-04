package models

import "time"

type Knowledge struct {
	ID          uint   `gorm:"primary_key"`
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text"`
	Type        string `gorm:"not null"`
	Content     string `gorm:"type:text;not null"`
	AuthorID    uint   `gorm:"not null"`
	Author      User   `gorm:"foreignKey:AuthorID"`
	Tags        []Tag  `gorm:"many2many:knowledge_tags;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
