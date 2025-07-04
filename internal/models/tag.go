package models

type Tag struct {
	ID         uint        `gorm:"primary_key"`
	Name       string      `gorm:"unique;not null"`
	Knowledges []Knowledge `gorm:"many2many:knowledge_tags;"`
}
